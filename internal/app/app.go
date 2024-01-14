package app

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ContentType int

const (
	Unsupported ContentType = iota
	PlainText
)

const ShortUrlBase = "http://localhost:8080"

type ContentTypes struct {
	name string
	code ContentType
}

var supportedTypes = []ContentTypes{
	{
		name: "text/plain",
		code: PlainText,
	},
}

// key is short url that corresponds to original url
var urls = make(map[string]string)

func checkContentType(name string) ContentType {
	for _, t := range supportedTypes {
		if name == t.name {
			return t.code
		}
	}
	return Unsupported
}

func generateShortKey() string {
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
func shortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if checkContentType(contentType) == PlainText {
			// we expect body contains url for encoding
			origUrl, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			shortKey := generateShortKey()
			urls[shortKey] = string(origUrl)
			shortenedUrl := fmt.Sprintf("%s/%s", ShortUrlBase, shortKey)
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
	} else if r.Method == http.MethodGet {
		// ignore
		if r.URL.RequestURI() == "/favicon.ico" {
		} else {
			shortKey := r.URL.RequestURI()[1:]
			origUrl, ok := urls[shortKey]
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

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", shortenURL)
	http.ListenAndServe("localhost:8080", mux)
}
