package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joshuabezaleel/unit-testing-golang-mocks/pkg/core/book"
)

const (
	serverPort = ":8083"
)

// Server holds dependencies for the HTTP server.
type Server struct {
	bookService book.Service

	Router *mux.Router
}

// NewServer returns a new HTTP server
// with all of the necessary dependencies.
func NewServer(bookService book.Service) *Server {
	server := &Server{
		bookService: bookService,
	}

	bookHandler := bookHandler{bookService}

	router := mux.NewRouter()

	bookHandler.registerRouter(router)

	server.Router = router

	return server
}

// Run runs the HTTP server with the port and the router.
func (srv *Server) Run() {
	err := http.ListenAndServe(serverPort, srv.Router)
	if err != nil {
		panic(err)
	}
}
