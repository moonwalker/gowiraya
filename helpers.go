package gowiraya

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *WirayaClient) apiPostOld(endpoint string, body interface{}, data interface{}) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest(http.MethodPost, c.baseUrlOldApi+endpoint, b)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-ApiKey", c.xApiKey)

	resp, err := c.HttpClientProxy.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *WirayaClient) apiCallNew(httpMethod string, endpoint string, body interface{}, data interface{}) error {
	bearerToken, err := c.getBearerToken()
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest(httpMethod, c.baseUrlNewApi+endpoint, b)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func (c *WirayaClient) getBearerToken() (string, error) {
	body := AuthRequestKey{
		Key: c.xApiKey,
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest(http.MethodPost, c.baseUrlNewApi+"/auth/token/apikey", b)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data := &TokenResponse{}

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return "", err
	}

	if data.Token == "" {
		return "", errors.New("token returned was empty")
	}

	return data.Token, nil
}
