package web

import (
	"net/http"
	"net/url"
	"strings"
	"wpc/pkg/data"
)

type Client struct {
	client *http.Client
}

func (c Client) Request(destinationUrl string, creds data.Creds) *http.Response {
	form := url.Values{}
	form.Add("log", creds.Username())
	form.Add("pwd", creds.Password())
	req, err := http.NewRequest("POST", destinationUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return &http.Response{StatusCode: -1}
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		return &http.Response{StatusCode: -1}
	}

	return resp
}

func NewClient() *Client {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &Client{client}
}
