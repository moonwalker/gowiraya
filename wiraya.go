package gowiraya

import (
	"crypto/tls"
	"fmt"
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
	HttpClientProxy *http.Client
	HttpClient      *http.Client
	baseUrlNewApi   string
	baseUrlOldApi   string

	xApiKey string
}

func NewWirayaClient(xApiKey string, proxy *string) (client *WirayaClient, err error) {
	httpClientProxy := http.DefaultClient
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

		httpClientProxy.Transport = &transport
	}

	client = &WirayaClient{
		HttpClientProxy: httpClientProxy,
		HttpClient:      httpClient,
		baseUrlOldApi:   "https://api.wiraya.com",
		baseUrlNewApi:   "https://api.wiraya.ai",
		xApiKey:         xApiKey,
	}

	return
}

//
// OLD api endpoints
//

func (c *WirayaClient) SendMessageFromAlpha(data SendMessage) (response Response, err error) {
	endpoint := "/api/SendMessageFromAlpha/json"

	err = c.apiPostOld(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) GetMessageStatus(data MessageStatus) (response Response, err error) {
	endpoint := "/api/GetMessageStatus/json"

	err = c.apiPostOld(endpoint, data, &response)
	if err != nil {
		return
	}

	return

}

func (c *WirayaClient) VerifyCode(data VerifyPinCode) (response Response, err error) {
	endpoint := "/api/VerifyCode/json"

	err = c.apiPostOld(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) SendPinCode(data SendPinCode) (response Response, err error) {
	endpoint := "/api/SendPinCode/json"

	err = c.apiPostOld(endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

//
// NEW api endpoints
//

func (c *WirayaClient) AddContact(contactID int64, data ContactRequstModel) (response ActionResponseWithId, err error) {
	endpoint := fmt.Sprintf("/api/Contact/%d", contactID)

	err = c.apiCallNew(http.MethodPut, endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) AddContactToCampaign(contactID int64, data CampaignRequestModel) (response IdResponse, err error) {
	endpoint := fmt.Sprintf("/api/Contact/%d/campaigns", contactID)

	err = c.apiCallNew(http.MethodPost, endpoint, data, &response)
	if err != nil {
		return
	}

	return
}

func (c *WirayaClient) AddEventToContact(contactID int64, data EventRequestModel) (response IdResponse, err error) {
	endpoint := fmt.Sprintf("/api/Contact/%d/events", contactID)

	err = c.apiCallNew(http.MethodPut, endpoint, data, &response)
	if err != nil {
		return
	}

	return
}
