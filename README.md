# GoWiraya

Go library for sending SMS with Wiraya.

Note that the ip needs to be whitelisted at Wiraya.

They also have some bugs
1. They do not support HTTP/2 over SSL
2. They do see headers as case sensitive (RESOLVED but not in production yet)

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