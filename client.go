package nclib

import (
	"context"

	"resty.dev/v3"
)

type Client interface {
	Close() error
	AddActivity(ctx context.Context, payload ...AddActivityPayload) error
}

type client struct {
	apiKey string
	client *resty.Client
	idc    IDC
}

func NewClient(idc IDC, opt ...Option) *client {
	rclient := resty.New()
	cl := client{
		// apiKey: apiKey,
		client: rclient,
		idc:    idc,
	}
	for _, o := range opt {
		o(&cl)
	}

	return &cl
}

func (c *client) WithApiKey(apiKey string) *client {
	newC := *c
	newC.apiKey = apiKey
	return &newC
}

func (c *client) req(ctx context.Context) *resty.Request {
	if c.client == nil {
		c.client = resty.New()
	}
	return c.client.R().SetAuthToken(c.apiKey).SetContext(ctx)
}

type Option func(*client)

func WithDebug(enable bool) Option {
	return func(c *client) {
		c.client = c.client.SetDebug(enable)
	}
}

func (c *client) Close() error {
	return c.client.Close()
}
