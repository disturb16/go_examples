package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/disturb16/go_examples/gorilla_mux/api"
	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"Hello World\"}")
}

func main() {

	r := mux.NewRouter()

	// create the api object
	a := &api.API{}

	// register the routes
	a.RegisterRoutes(r)

	r.HandleFunc("/", handleIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	log.Println("Listening...")
	srv.ListenAndServe()
}
