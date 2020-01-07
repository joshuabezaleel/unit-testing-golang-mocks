package book

// Repository provides access to the Book persistence store layer.
type Repository interface {
	Create(book *Book) (*Book, error)
	Get(bookID string) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(bookID string) error
}
