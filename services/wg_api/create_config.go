package wg_api

import (
	"XNetVPN-Back/config"
	"XNetVPN-Back/models/in"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func CreateWgConfig() (*in.ConfigResponse, error) {
	body := map[string]string{"InterfaceIdentifier": "wg0"}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", config.Config.WgServerApiUrl+"/provisioning/new-peer", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(config.Config.WgClientUsername, config.Config.WgClientPassword)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + resp.Status)
	}

	var configWg in.ConfigResponse
	err = json.NewDecoder(resp.Body).Decode(&configWg)
	if err != nil {
		return nil, err
	}

	return &configWg, nil
}
