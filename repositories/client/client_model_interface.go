package client

import "net/http"

type ClientModel interface {
	Get(url string) (*http.Response, error)
	Post(url string, body []byte) (*http.Response, error)
	Put(url string, body []byte) (*http.Response, error)
}
