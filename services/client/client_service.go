package client

import (
	"encoding/json"

	"github.com/dakasakti/todolist-web/deliveries/helpers"
	"github.com/dakasakti/todolist-web/repositories/client"
)

type clientService struct {
	cm client.ClientModel
}

func NewClientService(cm client.ClientModel) *clientService {
	return &clientService{cm: cm}
}

func (cs *clientService) GetData(url string) (helpers.ResponseJSON, error) {
	var result helpers.ResponseJSON
	resp, err := cs.cm.Get(url)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (cs *clientService) GetDatawithAuth(url string, token string) (helpers.ResponseJSON, error) {
	var result helpers.ResponseJSON
	resp, err := cs.cm.GetwithAuth(url, token)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (cs *clientService) Store(url string, reqBody []byte) (helpers.ResponseJSON, error) {
	var result helpers.ResponseJSON
	resp, err := cs.cm.Post(url, reqBody)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (cs *clientService) StorewithAuth(url string, token string, reqBody []byte) (helpers.ResponseJSON, error) {
	var result helpers.ResponseJSON
	resp, err := cs.cm.PostwithAuth(url, token, reqBody)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}

func (cs *clientService) UpdatewithAuth(url string, token string, reqBody []byte) (helpers.ResponseJSON, error) {
	var result helpers.ResponseJSON
	resp, err := cs.cm.PutwithAuth(url, token, reqBody)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&result)
	return result, nil
}
