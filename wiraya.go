package gowiraya

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
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

func NewWirayaClient(xApiKey string) (client *WirayaClient, err error) {
	// Disable HTTP/2 due to bug in Wiraya. They do not support HTTP/2 over SSL
	http.DefaultClient.Transport = &http.Transport{
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}

	client = &WirayaClient{
		HttpClient: http.DefaultClient,
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

func (c *WirayaClient) VerifyCode() {

}

func (c *WirayaClient) SendPinCode(data SendMessage) (response Response, err error){
	endpoint := "/api/SendPinCode/json"

	err = c.apiPost(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) AddCalListRecipient() {

}

func (c *WirayaClient) GetVoiceStatus() {

}

func (c *WirayaClient) apiPost(endpoint string, body interface{}, data interface{}) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(body)

	req, err := http.NewRequest(http.MethodPost, c.baseUrl+endpoint, b)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Adding header like this due to bug in Wiraya. They see headers as case sensitive..
	req.Header["X-ApiKey"] = []string{c.xApiKey}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}
