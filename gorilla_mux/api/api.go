package api

import (
	"fmt"
	"net/http"

	"github.com/disturb16/go_example/gorillamux/settings"
)

func New() *http.Server {

	port := settings.GetEnv("APP_PORT", "8080")
	addr := fmt.Sprintf(":%s", port)

	return &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handler),
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
