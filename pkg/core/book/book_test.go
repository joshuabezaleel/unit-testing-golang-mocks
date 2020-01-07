package book

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var bookRepository = &MockRepository{}
var bookService = service{bookRepository: bookRepository}

func TestCreate(t *testing.T) {
	book := NewBook("testID", "testTitle", "testAuthor")
	bookRepository.On("Create", book).Return(book, nil)

	createdBook, err := bookService.Create(book)

	require.Nil(t, err)
	require.Equal(t, createdBook.ID, book.ID)
}
