package tassemployee

import (
	"context"
	"net/url"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

type transport interface {
	request(context.Context, string, string, url.Values, any, ...int) ([]byte, error)
	upload(context.Context, string, tasscommon.FileRequest, ...int) ([]byte, error)
}

type Client struct {
	t transport
}

func NewClient(transport transport) Client {
	return Client{
		t: transport,
	}
}
