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
