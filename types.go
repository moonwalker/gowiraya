package gowiraya

type Response struct {
	Result string `json:"Result"`
}

type SendMessage struct {
	Recipient string `json:"recipient"` // Number to send pin to
	Sender    string `json:"sender"`    // Sender name for SMS
	Message   string `json:"message"`   // Message
}

type SendPinCode struct {
	Recipient string `json:"recipient"` // Number to send pin to
	Sender    string `json:"sender"`    // Sender name for SMS
	Message   string `json:"message"`   // Message, add {code} to specify where the pin should be
	IPAddress string `json:"ipaddress"` // IPAddress of user to avoid spam
}

type VerifyPinCode struct {
	Recipient string `json:"recipient"` // Number to send pin to
	IPAddress string `json:"ipaddress"` // IPAddress of user to avoid spam
	PinCode   string `json:"pincode"`   // Pin code that will be verified if active or not
}

type MessageStatus struct {
	MessageId string `json:"MessageId"`
}

type ContactRequstModel struct {
	Personal map[string]string `json:"personal"` // Data points with personally identifiable information
	General  map[string]string `json:"general"`  // Data points without PII
}

type ActionResponseWithId struct {
	Error    bool                   `json:"error,omitempty"`
	Failures map[string]interface{} `json:"failures,omitempty"`
}

type CampaignRequestModel struct {
	Campaign  string `json:"campaign"`            // Campaign identifier (provided by Wiraya)
	Iteration string `json:"iteration,omitempty"` // Iteration value. Only used with recurring communication to same contacts
}

type EventRequestModel struct {
	Name string `json:"name"`         // Event name (provided by Wiraya)
	At   string `json:"at,omitempty"` // Time of event. ISO format, in UTZ (Z)
}

type IdResponse struct {
	Id string `json:"id"` // Correlation id for ingested entity
}

type TokenResponse struct {
	Authenticated bool   `json:"authentication"`
	Token         string `json:"token"`
	TokenExpires  string `json:"tokenExpires"`
}

type AuthRequestKey struct {
	Key string `json:"key,omitempty"` // Static Wiraya API key
}
