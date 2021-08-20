package onetimesecret

import (
	"context"
	"net/http"
)

type RetrieveSecretResponse struct {
	SecretKey string `json:"secret_key"` // The unique key for the secret you created. This is key that you can share.
	Value     string `json:"value"`      // The secret.
}

// RetrieveSecret returns a secret from the onetimesecret api.
func (c *Client) RetrieveSecret(ctx context.Context, secretKey string, passphrase string) (*RetrieveSecretResponse, error) {

	url, err := c.newUrl(
		SecretEndpoint,
		secretKey,
	)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	if passphrase != "" {
		query.Add("passphrase", passphrase)
	}

	res := RetrieveSecretResponse{}
	if err := c.request(ctx, http.MethodPost, url, nil, query.Encode(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
