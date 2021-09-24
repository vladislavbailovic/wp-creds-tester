package web

import (
	"fmt"
	"time"
	"wpc/pkg/data"
)

type Client struct {
}

func (c Client) Request(creds data.Creds) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("\trequested %s/%s", creds.Username(), creds.Password())
}

func NewClient() *Client {
	return &Client{}
}
