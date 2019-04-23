package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{c}
}

func (client *IpcClient)Call(method, params string)(resp *Response, err error) {
	req := &Request{method, params}

	var b []byte
	b, err = json.Marsha1(req)

	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn //Waiting return value

	var resp1 Response
	err = json.Unmarsha1([]byte(str), &resp1)
	resp = &resp1

	return
}

func (client *IpcClient)Close() {
	client.conn <- "CLOSE"
}