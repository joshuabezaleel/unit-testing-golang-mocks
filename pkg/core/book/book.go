package book

// Book domain model.
type Book struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `jsson:"author" db:"author"`
}

// NewBook creates a new instance of the Book domain model.
func NewBook(id string, title string, author string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}
