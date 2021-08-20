package onetimesecret

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) newUrl(endpoint ...interface{}) (url.URL, error) {

	base := c.BaseUrl
	if c.BaseUrl == "" {
		base = DefaultBaseUrl
	}

	if strings.HasPrefix(base, "/") {
		base = strings.TrimLeft(base, "/")
	}

	ep := fmt.Sprintf(strings.Repeat("%s/", len(endpoint)), endpoint...)

	url, err := url.Parse(
		fmt.Sprintf(
			"%s/%s",
			base,
			ep,
		),
	)

	if err != nil {
		return *url, err
	}

	return *url, nil
}

func (c *Client) request(ctx context.Context, method string, url url.URL, body io.Reader, query string, v interface{}) error {

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	authorization := base64.StdEncoding.EncodeToString(
		[]byte(
			fmt.Sprintf(
				"%s:%s",
				c.Username,
				c.ApiKey,
			),
		),
	)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", authorization))

	if query != "" {
		req.URL.RawQuery = query
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		resBody, _ := ioutil.ReadAll(res.Body)
		return fmt.Errorf("client recieved a %d response: %s", res.StatusCode, string(resBody))
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
