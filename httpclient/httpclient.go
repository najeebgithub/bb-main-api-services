package httpclient

import (
	"net/http"
)

// Client is the set of methods for the http client
type HTTPClient interface {
	Request(request *Request) (*http.Response, error)
}

var client *Client

// Init is used to initialise the http client
func Init(configs ...*RequestConfig) {
	client = ConfigureHTTPClient(configs...)
}

// NewRequestConfig is used to create a new request config
func NewRequestConfig(name string, configs map[string]interface{}) *RequestConfig {
	return NewHTTPRequestConfig(name, configs)
}

// NewRequest is used to create a new request
func NewRequest(name string) *Request {
	return NewHTTPRequest(name)
}

// Get is used to get the client instance
func Get() HTTPClient {
	return client
}
