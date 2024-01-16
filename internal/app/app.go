package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/aleks0ps/url-service/internal/app/config"
	"github.com/aleks0ps/url-service/internal/app/handler"
)

func Run() {
	config.ParseOptions()
	r := chi.NewRouter()
	r.Get("/{id}", handler.GetOrigURL)
	r.Post("/", handler.ShortenURL)
	http.ListenAndServe(config.Options.ListenAddr, r)
}
