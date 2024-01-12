package main

import (
	"fmt"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, i love %s!", r.URL.Path[1:])
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", simpleHandler)
	http.ListenAndServe("localhost:8080", mux)
}
