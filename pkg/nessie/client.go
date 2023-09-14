package nessie

import (
	"net/http"
	"net/url"
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

func (c *Client) createURL(path string, queryArgs *map[string]string) (string, error) {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	base.Path += path

	params := url.Values{}
	if queryArgs != nil {
		for key, val := range *queryArgs {
			params.Add(key, val)
		}
	}
	params.Add("key", c.apiKey)
	base.RawQuery = params.Encode()
	return base.String(), nil
}
