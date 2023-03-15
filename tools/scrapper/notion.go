package scrapper

import (
	models "personal-backend/api/models/notion"
	"personal-backend/api/types"
	"personal-backend/utils"
)

func SynchronizeJournal() {
	utils.Logger().Log("SYNC", "Journal Synchronization")

	rows, err := models.NewJournalModelNotion().GetAllJournal()
	if err != nil {
		utils.Logger().Error(err)
		return
	}

	db := utils.DB()

	db.BeginTransaction()

	// _, err = db.Query(`delete from daily_journal`)
	// if err != nil {
	// 	db.Rollback()
	// 	utils.Logger().AddErrorData(err).Log("ENDSYNC", "Journal Synchronization Error, Transaction Rolled Back")
	// }

	for _, row := range rows {
		journalData := map[string]interface{}{
			"id":             row.ID,
			"title":          row.Title,
			"content":        row.Content,
			"date_created":   row.DateCreated,
			"date_modified":  row.DateModified,
			"notion_page_id": row.NotionPageID,
		}

		var notionID string

		err = db.QueryRow(`select notion_page_id from daily_journal where notion_page_id = $1`, journalData["notion_page_id"]).Scan(&notionID)
		if err != nil {
			db.Rollback()
			utils.Logger().AddErrorData(err).Log("ENDSYNC", "Journal Synchronization Error, Transaction Rolled Back")
			return
		}

		_, err = db.Insert(`daily_journal`, journalData)
		if err != nil {
			db.Rollback()
			utils.Logger().AddErrorData(err).Log("ENDSYNC", "Journal Synchronization Error, Transaction Rolled Back")
			return
		}
	}

	db.Rollback()
	utils.Logger().AddData(types.JournalData{Title: "123"}).Log("ENDSYNC", "Journal Synchronization Success")
}
