package onetimesecret

import (
	"context"
	"net/http"
)

type GetStatusResponse struct {
	Status string `json:"status"` //  The current system status. One of: nominal, offline.
}

// GetStatus returns the current status of the system.
func (c *Client) GetStatus(ctx context.Context) (*GetStatusResponse, error) {

	url, err := c.newUrl(
		c.BaseUrl,
		StatusEndpoint,
	)
	if err != nil {
		return nil, err
	}

	res := GetStatusResponse{}
	if err := c.request(ctx, http.MethodPost, url, nil, "", &res); err != nil {
		return nil, err
	}

	return &res, nil
}
