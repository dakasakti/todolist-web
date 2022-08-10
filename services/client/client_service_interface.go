package client

import "github.com/dakasakti/todolist-web/deliveries/helpers"

type ClientService interface {
	GetData(url string) (helpers.ResponseJSON, error)
	GetDatawithAuth(url string, token string) (helpers.ResponseJSON, error)
	Store(url string, reqBody []byte) (helpers.ResponseJSON, error)
	StorewithAuth(url string, token string, reqBody []byte) (helpers.ResponseJSON, error)
	UpdatewithAuth(url string, token string, reqBody []byte) (helpers.ResponseJSON, error)
}
