// Onetimesecret is a Go client that wraps the REST api of https://onetimesecret.com.
package onetimesecret

import (
	"net/http"
	"net/mail"
)

const (
	DefaultBaseUrl         = "https://onetimesecret.com/api/v1"
	StatusEndpoint         = "status"
	ShareEndpoint          = "share"
	GenerateEndpoint       = "generate"
	SecretEndpoint         = "secret"
	MetadataEndpoint       = "private"
	BurnEndpoint           = "burn"
	RecentMetadataEndpoint = "recent"
)

// HttpClient is used to abstract the requirement for http.Client
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	BaseUrl    string
	Username   string
	ApiKey     string
	HttpClient HttpClient
}

type ClientOption func(*Client)

func init() {}

// NewClient creates a new instance of the onetimesecret http client
func NewClient(options ...ClientOption) *Client {

	c := &Client{
		BaseUrl:    DefaultBaseUrl,
		Username:   "",
		ApiKey:     "",
		HttpClient: &http.Client{},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// WithBaseUrl overrides the clients default base url property
func WithBaseUrl(baseUrl string) ClientOption {
	return func(c *Client) {
		c.BaseUrl = baseUrl
	}
}

// WithUsername overrides the clients default username property
func WithUsername(username string) ClientOption {
	return func(c *Client) {

		_, err := mail.ParseAddress(username)
		if err != nil {
			panic(err)
		}

		c.Username = username
	}
}

// WithApiKey overrides the clients default api key property
func WithApiKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.ApiKey = apiKey
	}
}

// WithHttpClient overrides the default http client property
func WithHttpClient(client HttpClient) ClientOption {
	return func(c *Client) {
		c.HttpClient = client
	}
}
