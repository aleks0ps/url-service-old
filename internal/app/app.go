package app

import (
	"net/http"

	"github.com/aleks0ps/url-service/internal/app/handler"
)

func route(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handler.ShortenURL(w, r)
	case http.MethodGet:
		handler.GetOrigURL(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", route)
	http.ListenAndServe("localhost:8080", mux)
}
