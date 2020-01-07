package persistence

import (
	"github.com/jmoiron/sqlx"

	"github.com/joshuabezaleel/unit-testing-golang-mocks/pkg/core/book"
)

type bookRepository struct {
	DB *sqlx.DB
}

// NewBookRepository returns initialized implementations of the repository for
// Book domain model.
func NewBookRepository(DB *sqlx.DB) book.Repository {
	return &bookRepository{
		DB: DB,
	}
}

func (repo *bookRepository) Create(book *book.Book) (*book.Book, error) {
	_, err := repo.DB.NamedExec("INSERT INTO books (id, title, author) VALUES (:id, :title, :author)", book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (repo *bookRepository) Get(bookID string) (*book.Book, error) {
	book := book.Book{}
	err := repo.DB.QueryRowx("SELECT * FROM books WHERE id=$1", bookID).StructScan(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (repo *bookRepository) Update(book *book.Book) (*book.Book, error) {
	_, err := repo.DB.NamedExec("UDPATE books SET title=:title, author=:author WHERE id=:id", book)
	if err != nil {
		return nil, err
	}

	updatedBook, err := repo.Get(book.ID)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (repo *bookRepository) Delete(bookID string) error {
	_, err := repo.DB.Exec("DELETE FROM books WHERE id=$1", bookID)
	if err != nil {
		return err
	}

	return nil
}
