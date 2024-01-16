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
	UrlEncoded
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
		code: UrlEncoded,
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
		if checkContentType(contentType) == PlainText {
			// we expect body contains url for encoding
			origUrl, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			// empty body, no url supplied
			if len(origUrl) == 0 {
				w.WriteHeader(http.StatusBadRequest)
			}
			shortKey := GenerateShortKey()
			storage.StoreURL(shortKey, string(origUrl))
			shortenedUrl := fmt.Sprintf("%s/%s", config.Options.BaseURL, shortKey)
			// Return url
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(shortenedUrl)))
			// 201
			w.WriteHeader(http.StatusCreated)
			//
			fmt.Fprintf(w, shortenedUrl)
		} else if checkContentType(contentType) == UrlEncoded {
			r.ParseForm()
			origUrl := strings.Join(r.PostForm["url"], "")
			if len(origUrl) == 0 {
				w.WriteHeader(http.StatusBadRequest)
			}
			shortKey := GenerateShortKey()
			storage.StoreURL(shortKey, string(origUrl))
			shortenedUrl := fmt.Sprintf("%s/%s", config.Options.BaseURL, shortKey)
			// Return url
			w.Header().Set("Content-Type", "text/plain")
			w.Header().Set("Content-Length", strconv.Itoa(len(shortenedUrl)))
			// 201
			w.WriteHeader(http.StatusCreated)
			//
			fmt.Fprintf(w, shortenedUrl)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func GetOrigURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ignore
		if r.URL.RequestURI() == "/favicon.ico" {
		} else {
			shortKey := r.URL.RequestURI()[1:]
			origUrl, ok := storage.GetOrigURL(shortKey)
			if ok {
				http.Redirect(w, r, origUrl, http.StatusTemporaryRedirect)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
