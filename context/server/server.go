package server

import (
	"fmt"
	"net/http"
	"time"
)

func New() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-ctx.Done():
		fmt.Println("context cancelled ")
	case <-time.After(5 * time.Second):
		w.Write([]byte("Hello, World!"))
	}
}
