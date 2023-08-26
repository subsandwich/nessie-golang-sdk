package nessie

import (
	"fmt"
	"net/http"
)

type Option func(*Client)

func BaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func UnderlyingClient(client *http.Client) Option {
	return func(c *Client) {
		c.underlyingClient = client
	}
}

type Client struct {
	underlyingClient *http.Client
	baseURL          string
	apiKey           string
}

func New(apiKey string, opts ...Option) *Client {
	client := &Client{
		underlyingClient: &http.Client{},
		baseURL:          "http://api.nessieisreal.com",
		apiKey:           apiKey,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c *Client) createURL(path string) string {
	return fmt.Sprintf("%s/%s?key=%s", c.baseURL, path, c.apiKey)
}
