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

func Test_GetRecentMetadataReturnsAValidResponse(t *testing.T) {
	json := `{"custid":"test","metadata_key":"3Rg8R2sfD3?a", "secret_key": "xxxx", "ttl": 1, "metadata_ttl": 1, "secret_ttl": 1, "recipient": "test@test.com", "created": 1324174095, "updated": 1324174095, "received": 1324174095, "passphrase_required": true}`

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
	response, _ := client.GetRecentMetadata(ctx)

	assert.IsType(t, &[]GetRecentMetadataResponse{}, response)
}
