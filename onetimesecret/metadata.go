package onetimesecret

import (
	"context"
	"net/http"
)

// GetMetadataResponse represents the object returned from the api when requesting metadata for a secret.
type GetMetadataResponse struct {
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

// GetMetadata retrieves metadata about a secret.
// The metadata is intended to be used by the creator of the secret (i.e. not the recipient) and should generally be kept private.
// You can safely use the metadata key to retrieve basic information about the secret itself (e.g. if or when it was viewed) since the metadata key is different from the secret key.
func (c *Client) GetMetadata(ctx context.Context, metadataKey string) (*GetMetadataResponse, error) {

	url, err := c.newURL(
		MetadataEndpoint,
		metadataKey,
	)
	if err != nil {
		return nil, err
	}

	res := GetMetadataResponse{}
	if err := c.request(ctx, http.MethodPost, url, nil, "", &res); err != nil {
		return nil, err
	}

	return &res, nil
}
