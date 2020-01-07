package persistence

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")
	bookRepository := NewBookRepository(db)

	rows := sqlmock.NewRows([]string{"id, title, author"}).
		AddRow("testID", "testTitle", "testAuthor")

	mock.ExpectQuery("SELECT * FROM books WHERE id=?").
		WithArgs("testID").
		WillReturnRows(rows)

	newBook, err := bookRepository.Get("testID")
	require.Nil(t, err)
	require.Equal(t, newBook.ID, "testID")
}
