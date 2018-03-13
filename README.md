# GoWiraya

Go library for sending SMS with Wiraya.

Note that the ip needs to be whitelisted at Wiraya.

## Usage

```go
client, err := gowiraya.NewWirayaClient("xapikey")
if err != nil {
    fmt.Println(err)
    return
}

data := gowiraya.SendMessage{
    Recipient: "4612345678",
    Sender:    "Test",
    Message:   "This is the message",
}

resp, err := client.SendMessageFromAlpha(data)
if err != nil {
    fmt.Println(err)
    return
}

data2 := gowiraya.MessageStatus{
    MessageId: resp.Result,
}

resp2, err := client.GetMessageStatus(data2)
if err != nil {
    fmt.Println(err)
    return
}

fmt.Println("Status for message: " + resp2.Result)
```