package client

import "net/http"

type ClientModel interface {
	Get(url string) (*http.Response, error)
	GetwithAuth(url string, token string) (*http.Response, error)
	Post(url string, body []byte) (*http.Response, error)
	PostwithAuth(url string, token string, body []byte) (*http.Response, error)
	PutwithAuth(url string, token string, body []byte) (*http.Response, error)
}
