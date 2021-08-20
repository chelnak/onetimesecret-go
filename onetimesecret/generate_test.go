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
	passphrase = "test"
	ttlSeconds = 900
	recipient  = "test+test.com"
)

func Test_GenerateSecretReturnsAvalidResponse(t *testing.T) {

	json := `{"custid":"USERNAME","value":"3Rg8R2sfD3?a","metadata_key":"2b6bjmudhmtiqjn2qmdaqjkqxp323gi","secret_key":"pgcdv7org3vtdurif809sygnt0mstw6","ttl":"3600","updated":"1324174095","created":"1324174095"}`

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

	ctx := context.Background()
	response, _ := client.GenerateSecret(ctx, passphrase, ttlSeconds, recipient)

	assert.IsType(t, &GenerateSecretResponse{}, response)
}
