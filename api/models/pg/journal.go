package models

import (
	"fmt"
	"personal-backend/api/types"
	"personal-backend/utils"
)

type journalModel struct {
	types.PGModel
}

func NewJournalModel() types.IJournalModel {
	return &journalModel{
		types.PGModel{
			TableName: `daily_journal`,
		},
	}
}

func (m *journalModel) GetAllJournal() ([]types.JournalData, error) {
	var journalData []types.JournalData

	db := utils.DB()

	query := fmt.Sprintf(`select id, title, content, date_created, date_modified, notion_page_id from %s order by date_created desc`, m.TableName)
	rows, err := db.PreparedQuery(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var jd types.JournalData
		err = rows.Scan(
			&jd.ID,
			&jd.Title,
			&jd.Content,
			&jd.DateCreated,
			&jd.DateModified,
			&jd.NotionPageID,
		)
		if err != nil {
			return nil, err
		}
		journalData = append(journalData, jd)
	}

	return journalData, nil
}
