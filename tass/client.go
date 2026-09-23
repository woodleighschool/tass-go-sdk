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
	"strconv"
	"strings"
	"sync"
	"time"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
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
	AllowedCompanyCodes []TokenCompanyCode `json:"allowed_companies"`
}

func (r *TokenResponse) UnmarshalJSON(b []byte) error {
	var response struct {
		Token               string             `json:"token"`
		TokenExpiryDate     string             `json:"token_expiry_date"`
		AllowedCompanyCodes []TokenCompanyCode `json:"allowed_companies"`
	}

	if err := json.Unmarshal(b, &response); err != nil {
		return err
	}
	date, err := time.Parse("2006-01-02T03:04:05.000", response.TokenExpiryDate)
	if err != nil {
		return err
	}

	r.Token = response.Token
	r.TokenExpiryDate = date
	r.AllowedCompanyCodes = response.AllowedCompanyCodes

	return nil
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
	Modules      ModulesConfig
}

type ModulesConfig struct {
	Student  bool
	Employee bool
	Finance  bool
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
	apiEndpoint := strings.TrimRight(config.URL, "/") + "/api"
	return newClient(
		apiEndpoint,
		config.CompanyCode,
		config.Modules,
		config.ClientKey,
		config.ClientSecret,
		&http.Client{Timeout: defaultHTTPTimeout},
		time.Now,
	), nil
}

func newClient(
	apiEndpoint string,
	companyCode string,
	moduleConfig ModulesConfig,
	clientKey string,
	clientSecret string,
	httpClient *http.Client,
	now func() time.Time,
) *Client {
	client := Client{}

	client.transport = &transport{
		apiEndpoint:  strings.TrimRight(apiEndpoint, "/"),
		companyCode:  companyCode,
		clientKey:    clientKey,
		clientSecret: clientSecret,
		httpClient:   httpClient,
		now:          now,
	}

	if moduleConfig.Employee {
		c := tassemployee.NewClient(client.transport)
		client.Employee = &c
	}
	if moduleConfig.Student {
		c := tassstudent.NewClient(client.transport)
		client.Student = &c
	}
	if moduleConfig.Finance {
		c := tassfinance.NewClient(client.transport)
		client.Finance = &c
	}

	return &client
}

func (t *transport) Request(
	ctx context.Context,
	method string,
	path string,
	query url.Values,
	payload any,
	successCodes ...int,
) ([]byte, error) {
	encoded, err := encodeJSONPayload(payload)
	if err != nil {
		return nil, err
	}
	for attempt := range 2 {
		token, tokenErr := t.refreshToken(ctx)
		if tokenErr != nil {
			return nil, tokenErr
		}
		requestURL := fmt.Sprintf("%s/%s%s", t.apiEndpoint, t.companyCode, path)
		if len(query) != 0 {
			requestURL += "?" + query.Encode()
		}
		request, requestErr := http.NewRequestWithContext(ctx, method, requestURL, bytes.NewReader(encoded))
		if requestErr != nil {
			return nil, fmt.Errorf("create tass api request: %w", requestErr)
		}
		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, requestErr := t.httpClient.Do(request)
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
			t.invalidateToken(token)
			continue
		}
		if !containsStatus(successCodes, response.StatusCode) {
			return nil, NewHTTPError(requestURL, response.StatusCode, body)
		}
		return body, nil
	}
	return nil, fmt.Errorf("tass authentication failed after token refresh")
}

func (t *transport) Upload(
	ctx context.Context,
	path string,
	payload tasscommon.FileRequest,
	successCodes ...int,
) ([]byte, error) {
	encoded := encodeFormPayload(payload)
	for attempt := range 2 {
		token, tokenErr := t.refreshToken(ctx)
		if tokenErr != nil {
			return nil, tokenErr
		}
		requestURL := t.apiEndpoint + path
		request, requestErr := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, strings.NewReader(encoded))
		if requestErr != nil {
			return nil, fmt.Errorf("create tass api request: %w", requestErr)
		}
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Set("Content-Length", strconv.Itoa(len(encoded)))

		request.Header.Set("Accept", "application/json")
		request.Header.Set("Authorization", "Bearer "+token)

		response, requestErr := t.httpClient.Do(request)
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
			t.invalidateToken(token)
			continue
		}
		if !containsStatus(successCodes, response.StatusCode) {
			return nil, NewHTTPError(requestURL, response.StatusCode, body)
		}
		return body, nil
	}
	return nil, fmt.Errorf("tass authentication failed after token refresh")
}

func (t *transport) refreshToken(ctx context.Context) (string, error) {
	t.tokenMu.Lock()
	defer t.tokenMu.Unlock()

	if t.token != "" && t.now().Add(tokenRefreshLeeway).Before(t.tokenExpiryDate) {
		return t.token, nil
	}

	authPayload, err := json.Marshal(struct {
		ClientKey    string `json:"ClientKey"`
		ClientSecret string `json:"ClientSecret"`
	}{
		ClientKey:    t.clientKey,
		ClientSecret: t.clientSecret,
	})
	if err != nil {
		return "", fmt.Errorf("encode TASS token request: %w", err)
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		t.apiEndpoint+"/users",
		bytes.NewBuffer(authPayload),
	)
	if err != nil {
		return "", fmt.Errorf("create TASS API token request: %w", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	response, err := t.httpClient.Do(request)
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
		return "", fmt.Errorf("request TASS token: %w", NewHTTPError(t.apiEndpoint+"/users", response.StatusCode, body))
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
	if result.Token == "" || result.TokenExpiryDate.Before(t.now()) {
		return "", fmt.Errorf("decode TASS token: token and expiry in the future are required")
	}
	if !slices.ContainsFunc(result.AllowedCompanyCodes, func(allowedCompanyCode TokenCompanyCode) bool {
		return allowedCompanyCode.CompanyCode == t.companyCode
	}) {
		return "", fmt.Errorf("decode TASS token: company code %s is not in codes returned by API: %s", t.companyCode, result.AllowedCompanyCodes)
	}
	t.token = result.Token
	t.tokenExpiryDate = result.TokenExpiryDate
	return t.token, nil
}

func (t *transport) invalidateToken(token string) {
	t.tokenMu.Lock()
	defer t.tokenMu.Unlock()
	if t.token == token {
		t.token = ""
		t.tokenExpiryDate = time.Time{}
	}
}

func encodeJSONPayload(payload any) ([]byte, error) {
	if payload == nil {
		return nil, nil
	}
	encoded, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encode tass request body: %w", err)
	}
	return encoded, nil
}

func encodeFormPayload(payload tasscommon.FileRequest) string {
	data := url.Values{}

	if payload.FileName != nil {
		data.Set("file_name", *payload.FileName)
	}
	if len(payload.AdditionalProperties) != 0 {
		for key, value := range payload.AdditionalProperties {
			data.Set(key, value)
		}
	}

	data.Set("file_content", string(payload.FileContent))
	return data.Encode()
}

func containsStatus(statuses []int, status int) bool {
	return slices.Contains(statuses, status)
}
