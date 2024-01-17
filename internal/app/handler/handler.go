package handler

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aleks0ps/url-service/internal/app/config"
	"github.com/aleks0ps/url-service/internal/app/storage"
)

type ContentType int

const (
	Unsupported ContentType = iota
	PlainText
	URLEncoded
)

type ContentTypes struct {
	name string
	code ContentType
}

var supportedTypes = []ContentTypes{
	{
		name: "text/plain",
		code: PlainText,
	},
	{
		name: "application/x-www-form-urlencoded",
		code: URLEncoded,
	},
}

func checkContentType(name string) ContentType {
	for _, t := range supportedTypes {
		if name == t.name {
			return t.code
		}
	}
	return Unsupported
}

func GenerateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

// Send response to POST requests
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if checkContentType(contentType) == URLEncoded {
			r.ParseForm()
			origURL := strings.Join(r.PostForm["url"], "")
			// XXX
			if len(origURL) == 0 {
				w.WriteHeader(http.StatusBadRequest)
			}
			shortKey := GenerateShortKey()
			storage.StoreURL(shortKey, string(origURL))
			shortenedURL := fmt.Sprintf("%s/%s", config.Options.BaseURL, shortKey)
			// Return url
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
			// 201
			w.WriteHeader(http.StatusCreated)
			//
			fmt.Fprintf(w, shortenedURL)
		} else {
			origURL, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			shortKey := GenerateShortKey()
			storage.StoreURL(shortKey, string(origURL))
			shortenedURL := fmt.Sprintf("%s/%s", config.Options.BaseURL, shortKey)
			// Return url
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(shortenedURL)))
			// 201
			w.WriteHeader(http.StatusCreated)
			//
			fmt.Fprintf(w, shortenedURL)
		}
	}
}

func GetOrigURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ignore
		if r.URL.RequestURI() == "/favicon.ico" {
		} else {
			shortKey := r.URL.RequestURI()[1:]
			origURL, ok := storage.GetOrigURL(shortKey)
			if ok {
				http.Redirect(w, r, origURL, http.StatusTemporaryRedirect)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
