package book

// Service provides basic operations on Book domain model.
type Service interface {
	Create(book *Book) (*Book, error)
	Get(bookID string) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(bookID string) error
}

type service struct {
	bookRepository Repository
}

// NewBookService creates an instance of the service for the Book domain model
// with all of the necessary dependencies.
func NewBookService(bookRepository Repository) Service {
	return &service{
		bookRepository: bookRepository,
	}
}

func (s *service) Create(book *Book) (*Book, error) {
	return s.bookRepository.Create(book)
}

func (s *service) Get(bookID string) (*Book, error) {
	return s.bookRepository.Get(bookID)
}

func (s *service) Update(book *Book) (*Book, error) {
	return s.bookRepository.Update(book)
}

func (s *service) Delete(bookID string) error {
	return s.bookRepository.Delete(bookID)
}
