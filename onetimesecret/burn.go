package onetimesecret

import (
	"context"
	"net/http"
)

type BurnSecretResponse struct {
	CustId             string   `json:"custid"`              // The requestors user id
	MetadataKey        string   `json:"metadata_key"`        // The unique key for the metadata. DO NOT share this.
	SecretKey          string   `json:"secret_key"`          // The unique key for the secret you created. This is key that you can share.
	Ttl                int      `json:"ttl"`                 // The time-to-live that was specified (i.e. not the time remaining)
	MetadataTtl        int      `json:"metadata_ttl"`        // The remaining time (in seconds) that the metadata has left to live.
	SecretTtl          int      `json:"secret_ttl"`          // The remaining time (in seconds) that the secret has left to live.
	Recipient          []string `json:"recipient"`           // If a recipient was specified, this is an obfuscated version of the email address.
	Created            int64    `json:"created"`             // Time the metadata was created in unix time (UTC)
	Updated            int64    `json:"updated"`             // Time the metadata was last updated in unix time (UTC)
	PassphraseRequired bool     `json:"passphrase_required"` // If a passphrase was provided when the secret was created, this will be true. Otherwise false, obviously.
}

// BurnSecret will burn an unred secret.
func (c *Client) BurnSecret(ctx context.Context, metadataKey string) (*BurnSecretResponse, error) {

	url, err := c.newUrl(
		MetadataEndpoint,
		metadataKey,
		BurnEndpoint,
	)

	if err != nil {
		return nil, err
	}

	res := BurnSecretResponse{}
	if err := c.request(ctx, http.MethodPost, url, nil, "", &res); err != nil {
		return nil, err
	}

	return &res, nil
}
