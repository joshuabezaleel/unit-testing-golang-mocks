package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/joshuabezaleel/unit-testing-golang-mocks/pkg/core/book"
)

var bookService = &book.MockService{}
var bookTestingHandler = bookHandler{bookService}

func TestCreate(t *testing.T) {
	initialBook := book.NewBook("testID", "testTitle", "testAuthor")
	bookService.On("Create", initialBook).Return(initialBook, nil)

	url := fmt.Sprintf("/books")
	jsonReq, _ := json.Marshal(initialBook)
	req := httptest.NewRequest("POST", url, bytes.NewBuffer(jsonReq))
	rr := httptest.NewRecorder()

	bookTestingHandler.getBook(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	require.Nil(t, err)

	createdBook := book.Book{}
	err = json.Unmarshal(body, &createdBook)
	require.Nil(t, err)

	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, initialBook.ID, createdBook.ID)
}
