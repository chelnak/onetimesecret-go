// Package onetimesecret is a Go client that wraps the REST api of https://onetimesecret.com.
package onetimesecret

import (
	"net/http"
	"net/mail"
)

// Constants
const (
	DefaultBaseURL         = "https://onetimesecret.com/api/v1" // The default base url for the onetimesecret.com api.
	StatusEndpoint         = "status"                           // The status endpoint.
	ShareEndpoint          = "share"                            // The share endpoint.
	GenerateEndpoint       = "generate"                         // The generate endpoint.
	SecretEndpoint         = "secret"                           // The secret endpoint.
	MetadataEndpoint       = "private"                          // The private endpoint.
	BurnEndpoint           = "burn"                             // The burn endpoint.
	RecentMetadataEndpoint = "recent"                           // The recent metadata endpoint.
)

// HTTPClient is used to abstract the requirement for http.Client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client represents a onetimesecret http client instance.
type Client struct {
	BaseURL    string     // The base url of the api.
	Username   string     // The username of the requesting user.
	APIKey     string     // The api key of the requesting user.
	HTTPClient HTTPClient // A http.Client instance.
}

// ClientOption is the base struct for client options.
type ClientOption func(*Client)

func init() {}

// NewClient creates a new instance of the onetimesecret http client
func NewClient(options ...ClientOption) *Client {

	c := &Client{
		BaseURL:    DefaultBaseURL,
		Username:   "",
		APIKey:     "",
		HTTPClient: &http.Client{},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// WithBaseURL overrides the clients default base url property
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
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

// WithAPIKey overrides the clients default api key property
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.APIKey = apiKey
	}
}

// WithHTTPClient overrides the default http client property
func WithHTTPClient(client HTTPClient) ClientOption {
	return func(c *Client) {
		c.HTTPClient = client
	}
}
