package types

type JournalData struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	NotionPageID string `json:"notion_page_id"`
}

type IJournalModel interface {
	GetAllJournal() ([]JournalData, error)
}
