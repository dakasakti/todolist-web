package client

import (
	"bytes"
	"net/http"
)

type clientModel struct {
	client *http.Client
}

func NewClientModel() *clientModel {
	return &clientModel{
		client: &http.Client{},
	}
}

func (c *clientModel) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

func (c *clientModel) GetwithAuth(url string, token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return c.client.Do(req)
}

func (c *clientModel) Post(url string, body []byte) (*http.Response, error) {
	return c.client.Post(url, "application/json", bytes.NewBuffer(body))
}

func (c *clientModel) PostwithAuth(url string, token string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return c.client.Do(req)
}

func (c *clientModel) PutwithAuth(url string, token string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	return c.client.Do(req)
}
