package scrapper

import (
	models "personal-backend/api/models/notion"
	"personal-backend/utils"
)

func SynchronizeJournal() {
	utils.Logger().Log("SYNC", "Journal Synchronization")

	rows, err := models.NewJournalModel().GetAllJournal()
	if err != nil {
		utils.Logger().Error(err)
		return
	}

	db := utils.DB()

	db.BeginTransaction()

	newJournal := 0

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

		res, err := db.Query(`select notion_page_id from daily_journal where notion_page_id = $1`, journalData["notion_page_id"])
		if err != nil {
			db.Rollback()
			utils.Logger().AddErrorData(err).Log("ENDSYNC", "Journal Synchronization Error, Transaction Rolled Back")
			return
		}
		if res.Next() {
			res.Scan(&notionID)
			_, err := db.Update(`daily_journal`, map[string]interface{}{
				"title":         row.Title,
				"date_modified": row.DateModified,
				"date_created":  row.DateCreated,
				"content":       row.Content,
			}, map[string]interface{}{
				"notion_page_id": notionID,
			})
			if err != nil {
				db.Rollback()
				utils.Logger().AddErrorData(err).Log("ENDSYNC", "Journal Synchronization Error, Transaction Rolled Back")
				return
			}
		} else {
			_, err = db.Insert(`daily_journal`, journalData)
			if err != nil {
				db.Rollback()
				utils.Logger().AddErrorData(err).Log("ENDSYNC", "Journal Synchronization Error, Transaction Rolled Back")
				return
			}
			newJournal++
		}

	}

	db.Rollback()
	data := map[string]interface{}{
		"new_journal":    newJournal,
		"update_journal": len(rows) - newJournal,
	}
	utils.Logger().AddData(data).Log("ENDSYNC", "Journal Synchronization Success")
}
