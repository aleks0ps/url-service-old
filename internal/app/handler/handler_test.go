package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/aleks0ps/url-service/internal/app/storage"

	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	contentType := "text/plain"
	testCases := []struct {
		method       string
		body         string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodPost, body: "https://ya.ru", expectedCode: http.StatusCreated, expectedBody: ""},
		{method: http.MethodPost, body: "", expectedCode: http.StatusBadRequest, expectedBody: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", strings.NewReader(tc.body))
			r.Header.Set("Content-Type", contentType)
			w := httptest.NewRecorder()
			ShortenURL(w, r)
			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}

func TestGetOrigURL(t *testing.T) {
	contentType := "text/plain"
	baseUri := "http://localhost:8080/"
	urls := []struct {
		key     string
		origUrl string
	}{
		{key: GenerateShortKey(), origUrl: "https://ya.ru"},
		{key: GenerateShortKey(), origUrl: "https://google.com"},
	}

	for _, url := range urls {
		storage.StoreURL(url.key, url.origUrl)
	}

	testCases := []struct {
		method       string
		body         string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, body: baseUri + urls[0].key, expectedCode: http.StatusTemporaryRedirect, expectedBody: urls[0].origUrl},
		{method: http.MethodGet, body: baseUri + urls[1].key, expectedCode: http.StatusTemporaryRedirect, expectedBody: urls[1].origUrl},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, tc.body, nil)
			r.Header.Set("Content-Type", contentType)
			w := httptest.NewRecorder()
			GetOrigURL(w, r)
			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}
