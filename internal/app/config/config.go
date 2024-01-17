package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

var defaultAddress string = "localhost:8080"
var defaultBaseURL string = "http://localhost:8080"

var Options = struct {
	ListenAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
}{
	ListenAddr: defaultAddress,
	BaseURL:    defaultBaseURL,
}

func ParseOptions() {
	if err := env.Parse(&Options); err != nil {
	}
	flag.StringVar(&Options.ListenAddr, "a", Options.ListenAddr, "Listen address:port")
	flag.StringVar(&Options.BaseURL, "b", Options.BaseURL, "Base URL for shortened url")
	flag.Parse()
}
