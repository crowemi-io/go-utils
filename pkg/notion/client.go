package notion

import "github.com/dstotijn/go-notion"

type Client struct {
	client *notion.Client
}

func (c Client) CreateClient(apiKey string) {
	c.client = notion.NewClient(apiKey)
}
func (c Client) GetDatabase(databaseId string) (notion.Database, error) {
	var database notion.Database
	var err error
	return database, err
}
func (c Client) SearchDatabase() (notion.Database, error) {
	var database notion.Database
	var err error
	return database, err
}
func (c Client) GetPage(pageId string) (notion.Page, error) {
	var page notion.Page
	var err error
	return page, err
}
func (c Client) GetPages() ([]notion.Page, error) {
	var pages []notion.Page
	var err error
	return pages, err
}
