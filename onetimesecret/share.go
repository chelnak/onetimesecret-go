package onetimesecret

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/mail"
)

// ShareSecretResponse represents the object returned from the api when sharing a secret.
type ShareSecretResponse struct {
	CustID             string   `json:"custid"`              // The requestors user id
	MetadataKey        string   `json:"metadata_key"`        // The unique key for the metadata. DO NOT share this.
	SecretKey          string   `json:"secret_key"`          // The unique key for the secret you created. This is key that you can share.
	TTL                int      `json:"ttl"`                 // The time-to-live that was specified (i.e. not the time remaining)
	MetadataTTL        int      `json:"metadata_ttl"`        // The remaining time (in seconds) that the metadata has left to live.
	SecretTTL          int      `json:"secret_ttl"`          // The remaining time (in seconds) that the secret has left to live.
	Recipient          []string `json:"recipient"`           // If a recipient was specified, this is an obfuscated version of the email address.
	Created            int64    `json:"created"`             // Time the metadata was created in unix time (UTC)
	Updated            int64    `json:"updated"`             // Time the metadata was last updated in unix time (UTC)
	PassphraseRequired bool     `json:"passphrase_required"` // If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
}

// ShareSecret stores a secret value.
func (c *Client) ShareSecret(ctx context.Context, secret string, passphrase string, ttlSeconds int, recipient string) (*ShareSecretResponse, error) {

	url, err := c.newURL(
		c.BaseURL,
		ShareEndpoint,
	)
	if err != nil {
		return nil, err
	}

	if _, err := mail.ParseAddress(recipient); err != nil {
		return nil, errors.New("could not parse one or more recipient address")
	}

	query := url.Query()
	query.Add("secret", secret)
	query.Add("ttl", fmt.Sprint(ttlSeconds))
	query.Add("recipient", recipient)

	if passphrase != "" {
		query.Add("passphrase", passphrase)
	}

	res := ShareSecretResponse{}
	if err := c.request(ctx, http.MethodPost, url, nil, query.Encode(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
