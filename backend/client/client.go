package client

import (
	"net/http"
)

type Client struct {
	httpCli *http.Client
}

func NewClient(cli *http.Client) *Client {
	return &Client{httpCli: cli}
}

// func (cli *Client) check() error {
// 	return errors.Wrap(err error, message string)
// }
