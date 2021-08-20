package onetimesecret

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClientIsCreatedWithDefaults(t *testing.T) {

	client := NewClient()

	assert.Equal(t, DefaultBaseUrl, client.BaseUrl)
	assert.Equal(t, "", client.Username)
	assert.Equal(t, "", client.ApiKey)
}

func Test_ClientIsCreatedWithOptions(t *testing.T) {

	mockBaseUrl := "https://mytesturl.com"
	mockUsername := "testUser@gmail.com"
	mockApiKey := "testApiKey"

	client := NewClient(
		WithBaseUrl(mockBaseUrl),
		WithUsername(mockUsername),
		WithApiKey(mockApiKey),
	)

	assert.Equal(t, mockBaseUrl, client.BaseUrl, "")
	assert.Equal(t, mockUsername, client.Username, "")
	assert.Equal(t, mockApiKey, client.ApiKey, "")

}
