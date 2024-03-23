package main

import (
	"fmt"
	"net/http"
)

func getProfile(w http.ResponseWriter, req *http.Request) {
	profileID := req.PathValue("id")
	fmt.Fprint(w, "hello there, user", profileID)
}

func saveProfile(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "profile saved")
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /profiles/{id}", getProfile)
	r.HandleFunc("POST /profiles", saveProfile)
	http.ListenAndServe(":3000", r)
}
