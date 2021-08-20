package mocks

import "net/http"

var (
	// GetDoFunc acts as a setter when mocking with http.Client
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

// MockClient is used for mocking http.Client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do allows mocking http.Clients Do func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
