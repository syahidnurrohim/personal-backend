package utils

import (
	"context"
	"encoding/json"
	"os"

	"github.com/dstotijn/go-notion"
)

type NotionDatabase struct {
	databaseID string
	client     *notion.Client
}

type NotionPage struct {
	pageID string
	page   notion.Page
}

func NewNotionDatabase(databaseID string) *NotionDatabase {
	apiKey := os.Getenv("NOTION_SECRET_KEY")
	client := notion.NewClient(apiKey)
	return &NotionDatabase{databaseID: databaseID, client: client}
}

func (d *NotionDatabase) GetRows() ([]notion.Page, error) {
	res, err := d.client.QueryDatabase(
		context.Background(),
		d.databaseID,
		&notion.DatabaseQuery{})
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

func (d *NotionDatabase) GetBlockChildren(pageID string) ([]notion.Block, error) {
	res, err := d.client.FindBlockChildrenByID(
		context.Background(),
		pageID,
		&notion.PaginationQuery{})
	if err != nil {
		return nil, err
	}
	return res.Results, nil
}

func NewNotionPage(page notion.Page) *NotionPage {
	return &NotionPage{
		pageID: page.ID,
		page:   page,
	}
}

func (p *NotionPage) GetPageTitle() (string, error) {
	var prop notion.DatabasePageProperties
	mar, err := json.Marshal(p.page.Properties)
	if err != nil {
		return "", err
	}
	json.Unmarshal(mar, &prop)
	return prop["Name"].Title[0].PlainText, nil
}

func (p *NotionPage) GetPageCratedTime() string {
	return p.page.CreatedTime.Format("2006-01-02")
}

func (p *NotionPage) GetPageModifiedTime() string {
	return p.page.LastEditedTime.Format("2006-01-02")
}
