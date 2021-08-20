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

func Test_GetStatusReturnsAValidResponse(t *testing.T) {

	json := `{"status":"nominal"}`

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
	response, _ := client.GetStatus(ctx)

	assert.ObjectsAreEqual(GetStatusResponse{}, response)
	assert.Equal(t, "nominal", response.Status)
}
