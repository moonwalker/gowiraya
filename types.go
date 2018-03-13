package gowiraya

type Response struct {
	Result string `json:"Result"`
}

type SendMessage struct {
	Recipient string `json:"recipient"`           // Number to send pin to
	Sender    string `json:"sender"`              // Sender name for SMS
	Message   string `json:"message"`             // Message, add {code} to specify where the pin should be if pin request
	IPAddress string `json:"ipaddress,omitempty"` // IPAddress of user to avoid spam
}

type MessageStatus struct {
	MessageId string `json:"MessageId"`
}