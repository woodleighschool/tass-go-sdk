package tass

import (
	"errors"
	"fmt"
	"strings"
)

const MaxErrorBodyBytes = 4 << 10

type httpError struct {
	status int
	body   string
}

func (e *httpError) Error() string {
	if e.body == "" {
		return fmt.Sprintf("TASS API returned HTTP %d", e.status)
	}
	return fmt.Sprintf("TASS API returned HTTP %d: %s", e.status, e.body)
}

func NewHTTPError(status int, body []byte) error {
	if len(body) > MaxErrorBodyBytes {
		body = body[:MaxErrorBodyBytes]
	}
	return &httpError{status: status, body: strings.TrimSpace(string(body))}
}

func (c *Client) StatusCode(err error) (int, bool) {
	var apiError *httpError
	if !errors.As(err, &apiError) {
		return 0, false
	}
	return apiError.status, true
}
