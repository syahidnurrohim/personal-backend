package types

import (
	"personal-backend/utils"
)

type NotionModel struct {
	DatabaseID string
	Database   *utils.NotionDatabase
}

type PGModel struct {
	TableName string
}
