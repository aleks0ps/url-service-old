package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/aleks0ps/url-service/internal/app/handler"
)

func Run() {
	r := chi.NewRouter()
	r.Get("/{id}", handler.GetOrigURL)
	r.Post("/", handler.ShortenURL)
	http.ListenAndServe("localhost:8080", r)
}
