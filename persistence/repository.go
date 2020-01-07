package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Importing PostgreSQL driver

	"github.com/joshuabezaleel/unit-testing-golang-mocks/pkg/core/book"
)

const (
	dbHost     = "localhost"
	dbPort     = "8081"
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "unit-testing-golang-mocks"
	dbDriver   = "postgres"
)

// Repository holds dependencies for the current persistence layer.
type Repository struct {
	BookRepository book.Repository

	DB *sqlx.DB
}

// NewRepository returns a new Repository
// with all of the necessary dependencies.
func NewRepository() *Repository {
	connectionString := fmt.Sprintf("host = %s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sqlx.Open(dbDriver, connectionString)
	if err != nil {
		panic(err)
	}

	bookRepository := NewBookRepository(db)

	repository := &Repository{
		BookRepository: bookRepository,
		DB:             db,
	}

	return repository
}
