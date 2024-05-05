package main

import (
	"fmt"
	"net/http"
)

func handleGetUser(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello there "+id)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /users/{id}", handleGetUser)

	http.ListenAndServe(":8080", r)
}
