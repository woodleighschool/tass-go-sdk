package tassstudent

import (
	"context"
	"net/url"
)

type transport interface {
	request(context.Context, string, string, url.Values, any, ...int) ([]byte, error)
}

type Client struct {
	t transport
}

func NewClient(transport transport) Client {
	return Client{
		t: transport,
	}
}
