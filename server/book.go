package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joshuabezaleel/unit-testing-golang-mocks/pkg/core/book"
)

type bookHandler struct {
	bookService book.Service
}

func (handler *bookHandler) registerRouter(router *mux.Router) {
	router.HandleFunc("/books", handler.createBook).Methods("POST")
	router.HandleFunc("/books/{bookID}", handler.getBook).Methods("GET")
	router.HandleFunc("/books/{bookID}", handler.updateBook).Methods("PUT")
	router.HandleFunc("/books/{bookID}", handler.deleteBook).Methods("DELETE")
}

func (handler *bookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	book := book.Book{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	newBook, err := handler.bookService.Create(&book)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, newBook)
}

func (handler *bookHandler) getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := handler.bookService.Get(bookID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func (handler *bookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	book := book.Book{}

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	book.ID = bookID

	updatedBook, err := handler.bookService.Update(&book)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, updatedBook)
}

func (handler *bookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, ok := vars["bookID"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	err := handler.bookService.Delete(bookID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, "Book "+bookID+" deleted")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"Error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
