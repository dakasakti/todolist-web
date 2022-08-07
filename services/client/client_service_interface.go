package client

import "github.com/dakasakti/todolist-web/deliveries/helpers"

type ClientService interface {
	GetData(url string) (helpers.ResponseJSON, error)
	Store(url string, reqBody []byte) error
	Update(url string, reqBody []byte) error
}
