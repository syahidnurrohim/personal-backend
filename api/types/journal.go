package types

type JournalData struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Content     []string `json:"content"`
	DateCreated string   `json:"date_created"`
}

type IJournalModel interface {
	GetAllJournal() ([]JournalData, error)
}
