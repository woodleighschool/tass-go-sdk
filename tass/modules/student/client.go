package tassstudent

import (
	"context"
	"net/url"

	tasscommon "github.com/woodleighschool/tass-go-sdk/tass/modules/common"
)

type transport interface {
	Request(context.Context, string, string, url.Values, any, ...int) ([]byte, error)
	Upload(context.Context, string, tasscommon.FileRequest, ...int) ([]byte, error)
}

type Client struct {
	t transport
}

func NewClient(transport transport) Client {
	return Client{
		t: transport,
	}
}
