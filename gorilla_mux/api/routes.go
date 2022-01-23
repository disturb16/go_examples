package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Router) {
	r.Use(requestIDHandler)

	public := r.NewRoute().Subrouter()
	protected := r.NewRoute().Subrouter()

	protected.Use(authMiddleware)

	public.HandleFunc("/books", a.getBooks).Methods(http.MethodGet)
	public.HandleFunc("/books/{id}", a.getBook).Methods(http.MethodGet)

	protected.HandleFunc("/books", a.postBook).Methods(http.MethodPost)
}
