package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/books", a.getBooks).Methods(http.MethodGet)
	r.HandleFunc("/books", a.postBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", a.getBook).Methods(http.MethodGet)
}
