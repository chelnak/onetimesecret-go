package onetimesecret

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ClientIsCreatedWithDefaults(t *testing.T) {

	client := NewClient()

	assert.Equal(t, DefaultBaseURL, client.BaseURL)
	assert.Equal(t, "", client.Username)
	assert.Equal(t, "", client.APIKey)
}

func Test_ClientIsCreatedWithOptions(t *testing.T) {

	mockBaseURL := "https://mytesturl.com"
	mockUsername := "testUser@gmail.com"
	mockAPIKey := "testApiKey"

	client := NewClient(
		WithBaseURL(mockBaseURL),
		WithUsername(mockUsername),
		WithAPIKey(mockAPIKey),
	)

	assert.Equal(t, mockBaseURL, client.BaseURL, "")
	assert.Equal(t, mockUsername, client.Username, "")
	assert.Equal(t, mockAPIKey, client.APIKey, "")

}
