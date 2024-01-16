package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

var Options struct {
	ListenAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
}

var defaultAddress string = "localhost:8080"
var defaultBaseURL string = "http://localhost:8080"

func ParseOptions() {
	if err := env.Parse(&Options); err != nil {
		Options.ListenAddr = defaultAddress
		Options.BaseURL = defaultBaseURL
	}
	flag.StringVar(&Options.ListenAddr, "a", defaultAddress, "Listen address:port")
	flag.StringVar(&Options.BaseURL, "b", defaultBaseURL, "Base URL for shortened url")
	flag.Parse()
}
