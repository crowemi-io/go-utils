package notion

import "github.com/dstotijn/go-notion"

type Client struct {
	client *notion.Client
}

func (c Client) CreateClient(apiKey string) {
	c.client = notion.NewClient(apiKey)
}
