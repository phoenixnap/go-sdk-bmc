package client

import (
	"io"
	"net/http"
)

// Requester is the interface used to represent a Client that performs requests.
type Requester interface {
	Get(resource string) (*http.Response, error)
	Post(resource string, body io.Reader) (*http.Response, error)
	Delete(resource string) (*http.Response, error)
}
