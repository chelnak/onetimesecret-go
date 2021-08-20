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

const (
	secretKey = "test"
)

func Test_RetrieveSecretReturnsAValidResponse(t *testing.T) {
	json := `{"secret_key":"test","value":"3Rg8R2sfD3?a"}`

	mockClient := &mocks.MockClient{}
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
		}, nil
	}

	client := NewClient(
		WithHttpClient(mockClient),
	)

	ctx := context.Background()
	response, _ := client.RetrieveSecret(ctx, secretKey, passphrase)

	assert.IsType(t, &RetrieveSecretResponse{}, response)
}
