package main

import (
	"github.com/joshuabezaleel/unit-testing-golang-mocks/persistence"
	"github.com/joshuabezaleel/unit-testing-golang-mocks/pkg/core/book"
	"github.com/joshuabezaleel/unit-testing-golang-mocks/server"
)

func main() {
	repository := persistence.NewRepository()

	bookService := book.NewBookService(repository.BookRepository)

	srv := server.NewServer(bookService)
	srv.Run()
}
