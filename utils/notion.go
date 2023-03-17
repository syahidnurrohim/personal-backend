package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
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

type NotionBlock struct {
	block notion.Block
}

func NewNotionDatabase(databaseID string) *NotionDatabase {
	apiKey := os.Getenv("NOTION_SECRET_KEY")
	spew.Dump(apiKey)
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

func NewNotionBlock(block notion.Block) *NotionBlock {
	return &NotionBlock{
		block: block,
	}
}

func (b *NotionBlock) GetBlockPlainText() (string, error) {
	mar, err := b.block.MarshalJSON()
	if err != nil {
		return "", err
	}

	blockMap := make(map[string]interface{})

	err = json.Unmarshal(mar, &blockMap)
	if err != nil {
		return "", err
	}

	var findPlainText func(blockMap map[string]interface{}) []string

	findPlainText = func(blockMap map[string]interface{}) []string {
		plainText := []string{}
		for k, v := range blockMap {
			if k == "plain_text" {
				plainText = append(plainText, fmt.Sprintf("%s", v))
				continue
			}

			blockMap2, err := ToMap(v)
			if err == nil {
				plainText = append(plainText, findPlainText(blockMap2)...)
				continue
			}

			blockSlice, err := ToSlice(v)
			if err == nil {
				for _, bs := range blockSlice {
					blockMap2, err := ToMap(bs)
					if err == nil {
						plainText = append(plainText, findPlainText(blockMap2)...)
					}
				}
			}

		}
		return plainText
	}

	joinedText := strings.Join(findPlainText(blockMap), "<br>")

	return joinedText, nil
}
