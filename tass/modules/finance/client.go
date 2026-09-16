package tassfinance

import (
	"context"
	"net/url"
)

type transport interface {
	request(context.Context, string, string, url.Values, ...int)
}

type Client struct {
	transport transport
}

func NewClient(transport transport) Client {
	return Client{
		transport: transport,
	}
}
