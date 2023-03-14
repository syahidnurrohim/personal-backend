package models

import (
	"encoding/json"
	"fmt"
	"personal-backend/api/types"
	"personal-backend/utils"

	"github.com/davecgh/go-spew/spew"
	"github.com/dstotijn/go-notion"
)

type journalModel struct {
	types.NotionModel
}

type journalParagraph struct {
	RichText []notion.RichText `json:"rich_text"`
	Color    string            `json:"color"`
}

type journalBlock struct {
	Paragraph journalParagraph `json:"paragraph"`
}

func NewJournalModel() types.IJournalModel {
	databaseID := "026f69c6d7d64555a893a8218185dd8b"
	database := utils.NewNotionDatabase(databaseID)
	return &journalModel{
		types.NotionModel{
			DatabaseID: databaseID,
			Database:   database,
		},
	}
}

func (m *journalModel) GetAllJournal() ([]types.JournalData, error) {
	rows, err := m.Database.GetRows()
	if err != nil {
		return nil, err
	}

	var journalData []types.JournalData

	for _, row := range rows {
		page := utils.NewNotionPage(row)

		title, err := page.GetPageTitle()
		if err != nil {
			return nil, err
		}

		blockChildren, err := m.Database.GetBlockChildren(row.ID)
		if err != nil {
			return nil, fmt.Errorf("page id %s not found", row.ID)
		}

		spew.Dump(blockChildren)
		var content []string

		for _, c := range blockChildren {
			var block journalBlock

			mar, err := c.MarshalJSON()
			if err != nil {
				return nil, fmt.Errorf("cannot marshal block id %s to json", c.ID())
			}

			if json.Unmarshal(mar, &block) != nil {
				return nil, fmt.Errorf("cannot unmarshal block id %s to journal block type", c.ID())
			}

			for _, b := range block.Paragraph.RichText {
				content = append(content, b.PlainText)
			}
		}

		journalData = append(journalData, types.JournalData{
			ID:          row.ID,
			Title:       title,
			Content:     content,
			DateCreated: page.GetPageCratedTime(),
		})
		break
	}
	return journalData, nil
}
