package nclib

import (
	"resty.dev/v3"
)

type client struct {
	apiKey string
	client *resty.Client
	idc    IDC
}

func NewClient(apiKey string, idc IDC, opt ...Option) *client {
	rclient := resty.New().SetAuthToken(apiKey)
	cl := client{
		apiKey: apiKey,
		client: rclient,
		idc:    idc,
	}
	for _, o := range opt {
		o(&cl)
	}

	return &cl
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
