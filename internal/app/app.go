package app

import (
	"net/http"

	"github.com/aleks0ps/url-service/internal/app/handler"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.ShortenURL)
	mux.HandleFunc("/{id}", handler.GetOrigURL)
	http.ListenAndServe("localhost:8080", mux)
}
