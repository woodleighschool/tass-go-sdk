package tass

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"

	tassemployee "github.com/woodleighschool/tass-go-sdk/tass/modules/employee"
	tassfinance "github.com/woodleighschool/tass-go-sdk/tass/modules/finance"
	tassstudent "github.com/woodleighschool/tass-go-sdk/tass/modules/student"
)

const (
	defaultHTTPTimeout = 30 * time.Second
	tokenRefreshLeeway = 30 * time.Second
)

type TokenResponse struct {
	Token               string             `json:"token"`
	TokenExpiryDate     time.Time          `json:"token_expiry_date"`
	AllowedCompanyCodes []TokenCompanyCode `json:"allowed_company_codes"`
}

type TokenCompanyCode struct {
	CompanyCode string `json:"cmpy_code"`
	Description string `json:"name_text"`
}

type Config struct {
	URL          string
	CompanyCode  string
	ClientKey    string
	ClientSecret string
	CommTypeCode string
}

type Client struct {
	transport *transport
	Student   *tassstudent.Client
	Employee  *tassemployee.Client
	Finance   *tassfinance.Client
}

type transport struct {
	apiEndpoint  string
	companyCode  string
	commTypeCode string
	clientKey    string
	clientSecret string
	httpClient   *http.Client
	now          func() time.Time

	tokenMu         sync.Mutex
	token           string
	tokenExpiryDate time.Time
}

func NewClient(c any) (*Client, error) {
	config, ok := c.(Config)
	if !ok {
		return nil, fmt.Errorf("unable to assert config as tass config: %s", c)
	}
	if strings.TrimSpace(config.ClientKey) == "" || config.ClientSecret == "" {
		return nil, fmt.Errorf("tass Client ID and secret are required")
	}
	parsed, err := url.Parse(config.URL)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" {
		return nil, fmt.Errorf("a TASS host is required: https://{your_host}")
	}
	if parsed.RawQuery != "" || parsed.Fragment != "" {
		return nil, fmt.Errorf("tass URL: must not contain query or fragment")
	}
	validCompanyCode, err := regexp.MatchString("^[0-9]{2}$", config.CompanyCode)
	if err != nil || !validCompanyCode {
		return nil, fmt.Errorf("tass company code must be a 2 digit number")
	}
	apiEndpoint := strings.TrimRight(config.URL, "/") + "/api/" + config.CompanyCode
	return newClient(
		apiEndpoint,
		config.CompanyCode,
		config.CommTypeCode,
		config.ClientKey,
		config.ClientSecret,
		&http.Client{Timeout: defaultHTTPTimeout},
		time.Now,
	), nil
}

func newClient(
	apiEndpoint string,
	companyCode string,
	commTypeCode string,
	clientKey string,
	clientSecret string,
	httpClient *http.Client,
	now func() time.Time,
) *Client {
	return &Client{
		transport: &transport{
			apiEndpoint:  strings.TrimRight(apiEndpoint, "/"),
			companyCode:  companyCode,
			commTypeCode: commTypeCode,
			clientKey:    clientKey,
			clientSecret: clientSecret,
			httpClient:   httpClient,
			now:          now,
		},
	}
}

func (c *Client) request(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	payload any,
	successCodes ...int,
) ([]byte, error) {
	encoded, err := encodePayload(payload)
	if err != nil {
		return nil, err
	}
	for attempt := range 2 {
		token, tokenErr := c.refreshToken(ctx)
		if tokenErr != nil {
			return nil, tokenErr
		}
		requestURL := c.transport.apiEndpoint + path
		if len(query) != 0 {
			requestURL += "?" + query.Encode()
		}
		request, requestErr := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewReader(encoded))
		if requestErr != nil {
			return nil, fmt.Errorf("create tass api request: %w", requestErr)
		}
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, requestErr := c.transport.httpClient.Do(request)
		if requestErr != nil {
			return nil, fmt.Errorf("call tass api: %w", requestErr)
		}
		body, readErr := io.ReadAll(response.Body)
		closeErr := response.Body.Close()
		if readErr != nil {
			return nil, fmt.Errorf("read tass response: %w", readErr)
		}
		if closeErr != nil {
			return nil, fmt.Errorf("close tass response: %w", closeErr)
		}
		if response.StatusCode == http.StatusUnauthorized && attempt == 0 {
			c.invalidateToken(token)
			continue
		}
		if !containsStatus(successCodes, response.StatusCode) {
			return nil, NewHTTPError(response.StatusCode, body)
		}
		return body, nil
	}
	return nil, fmt.Errorf("tass authentication failed after token refresh")
}

func (c *Client) refreshToken(ctx context.Context) (string, error) {
	c.transport.tokenMu.Lock()
	defer c.transport.tokenMu.Unlock()

	if c.transport.token != "" && c.transport.now().Add(tokenRefreshLeeway).Before(c.transport.tokenExpiryDate) {
		return c.transport.token, nil
	}

	authPayload, err := json.Marshal(struct {
		ClientKey    string `json:"ClientKey"`
		ClientSecret string `json:"ClientSecret"`
	}{
		ClientKey:    c.transport.clientKey,
		ClientSecret: c.transport.clientSecret,
	})
	if err != nil {
		return "", fmt.Errorf("encode TASS token request: %w", err)
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.transport.apiEndpoint+"/users",
		bytes.NewBuffer(authPayload),
	)
	if err != nil {
		return "", fmt.Errorf("create TASS API token request: %w", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	response, err := c.transport.httpClient.Do(request)
	if err != nil {
		return "", fmt.Errorf("request TASS token: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		body, readErr := io.ReadAll(io.LimitReader(response.Body, MaxErrorBodyBytes+1))
		closeErr := response.Body.Close()
		if readErr != nil {
			return "", fmt.Errorf("read TASS token error: %w", readErr)
		}
		if closeErr != nil {
			return "", fmt.Errorf("close TASS token error response: %w", closeErr)
		}
		return "", fmt.Errorf("request TASS token: %w", NewHTTPError(response.StatusCode, body))
	}
	var result TokenResponse
	decodeErr := json.NewDecoder(response.Body).Decode(&result)
	closeErr := response.Body.Close()
	if decodeErr != nil {
		return "", fmt.Errorf("decode TASS token: %w", decodeErr)
	}
	if closeErr != nil {
		return "", fmt.Errorf("close TASS token response: %w", closeErr)
	}
	if result.Token == "" || result.TokenExpiryDate.Before(c.transport.now()) {
		return "", fmt.Errorf("decode TASS token: token and expiry in the future are required")
	}
	if !slices.ContainsFunc(result.AllowedCompanyCodes, func(allowedCompanyCode TokenCompanyCode) bool {
		return allowedCompanyCode.CompanyCode == c.transport.companyCode
	}) {
		return "", fmt.Errorf("decode TASS token: company code %s is not in codes returned by API: %s", c.transport.companyCode, result.AllowedCompanyCodes)
	}
	c.transport.token = result.Token
	c.transport.tokenExpiryDate = result.TokenExpiryDate
	return c.transport.token, nil
}

func (c *Client) invalidateToken(token string) {
	c.transport.tokenMu.Lock()
	defer c.transport.tokenMu.Unlock()
	if c.transport.token == token {
		c.transport.token = ""
		c.transport.tokenExpiryDate = time.Time{}
	}
}

func encodePayload(payload any) ([]byte, error) {
	if payload == nil {
		return nil, nil
	}
	encoded, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encode tass request body: %w", err)
	}
	return encoded, nil
}

func containsStatus(statuses []int, status int) bool {
	return slices.Contains(statuses, status)
}
