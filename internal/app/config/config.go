package config

import "flag"

var Options struct {
	ListenAddr string
	BaseURL    string
}

var defaultAddress string = "localhost:8080"
var defaultBaseURL string = "http://localhost:8080"

func ParseOptions() {
	flag.StringVar(&Options.ListenAddr, "a", defaultAddress, "Listen address:port")
	flag.StringVar(&Options.BaseURL, "b", defaultBaseURL, "Base URL for shortened url")
	flag.Parse()
}
