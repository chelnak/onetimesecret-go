//+build unit

package onetimesecret

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/chelnak/onetimesecret-go/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_ShareSecretReturnsAValidResponse(t *testing.T) {

	json := `{"custid":"USERNAME","metadata_key":"qjpjroeit8wra0ojeyhcw5pjsgwtuq7","secret_key":"153l8vbwqx5taskp92pf05uvgjefvu9","ttl":"3600","updated":"1324174006","created":"1324174006"}`

	mockClient := &mocks.MockClient{}
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
		}, nil
	}

	client := NewClient(
		WithHTTPClient(mockClient),
	)

	secret := "testSecret"
	passphrase := "test"
	ttlSeconds := 900
	recipient := "test@test.com"

	ctx := context.Background()
	response, _ := client.ShareSecret(ctx, secret, passphrase, ttlSeconds, recipient)
	assert.IsType(t, &ShareSecretResponse{}, response)
	//assert.Equal(t, "USERNAME", response.CustId)
}

func Test_ShareSecretReturnsAnErrorWhenAnInvalidRecipientIsPassed(t *testing.T) {

	json := `{"custid":"USERNAME","metadata_key":"qjpjroeit8wra0ojeyhcw5pjsgwtuq7","secret_key":"153l8vbwqx5taskp92pf05uvgjefvu9","ttl":"3600","updated":"1324174006","created":"1324174006"}`

	mockClient := &mocks.MockClient{}
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
		}, nil
	}

	client := NewClient(
		WithHTTPClient(mockClient),
	)

	secret := "testSecret"
	passphrase := "test"
	ttlSeconds := 900
	recipient := "test+test.com"

	ctx := context.Background()
	_, err := client.ShareSecret(ctx, secret, passphrase, ttlSeconds, recipient)
	assert.Error(t, err)
}
