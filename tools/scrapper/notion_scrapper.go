package scrapper

import "personal-backend/utils"

func SynchronizeJournal() {
	databaseID := "026f69c6d7d64555a893a8218185dd8b"
	db := utils.NewNotionDatabase(databaseID)
	db.GetRows()
}
