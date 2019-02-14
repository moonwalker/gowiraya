package gowiraya

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	ErrorSender          = "ERROR_SENDER"
	ErrorContent         = "ERROR_CONTENT"
	ErrorDateTime        = "ERROR_DATETIME"
	ErrorAuthentication  = "ERROR_AUTHENTICATION"
	ErrorProjectNotFound = "ERROR_PROJECT_NOT_FOUND"

	StatusInqueue      = "INQUEUE"
	StatusSent         = "SENT"
	StatusDelivered    = "DELIVERED"
	StatusNotDelivered = "NOT DELIVERED"
	StatusTooLate      = "TOOLATE"
	StatusFailure      = "FAILURE"
)

type WirayaClient struct {
	HttpClient *http.Client
	baseUrl    string
	xApiKey    string
}

func NewWirayaClient(xApiKey string, proxy *string) (client *WirayaClient, err error) {
	httpClient := http.DefaultClient

	if proxy != nil {
		proxyURL, err := url.Parse(*proxy)
		if err != nil {
			return nil, err
		}

		transport := http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{},
		}

		httpClient.Transport = &transport
	}

	client = &WirayaClient{
		HttpClient: httpClient,
		baseUrl:    "https://api.wiraya.com",
		xApiKey:    xApiKey,
	}

	return
}

func (c *WirayaClient) SendMessageFromAlpha(data SendMessage) (response Response, err error) {
	endpoint := "/api/SendMessageFromAlpha/json"

	err = c.apiPost(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) GetMessageStatus(data MessageStatus) (response Response, err error) {
	endpoint := "/api/GetMessageStatus/json"

	err = c.apiPost(endpoint, data, &response)
	if err != nil {
		return
	}

	return

}

func (c *WirayaClient) VerifyCode(data VerifyPinCode) (response Response, err error) {
	endpoint := "/api/VerifyCode/json"

	err = c.apiPost(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) SendPinCode(data SendPinCode) (response Response, err error) {
	endpoint := "/api/SendPinCode/json"

	err = c.apiPost(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) apiPost(endpoint string, body interface{}, data interface{}) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest(http.MethodPost, c.baseUrl+endpoint, b)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-ApiKey", c.xApiKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

/*

Not yet implemented api endpoints

func (c *WirayaClient) CreateSMSProject() {

}

func (c *WirayaClient) GetSMSProjectInfo() {

}

func (c *WirayaClient) UpdateSMSProject() {

}

func (c *WirayaClient) AddSMSRecipients() {

}

func (c *WirayaClient) AddSMSRecipient() {

}

func (c *WirayaClient) AddCalListRecipient() {

}

func (c *WirayaClient) GetVoiceStatus() {

}
*/