package config

import (
	"flag"
	"os"
	"strings"
)

type Options struct {
	ServerURL string
	ShortURL  string
}

func Get() *Options {
	options := new(Options)
	options.ShortURL = getShURL()
	options.ServerURL = getSrvURL()

	validateOptions(options)

	return options
}

func getShURL() string {
	s, f := os.LookupEnv("SERVER_ADDRESS")
	if f && s != "" {
		return s
	}

	flag.StringVar(&s, "b", "http://localhost:8080/", "базовый адрес результирующего шортлинка")
	flag.Parse()

	return s
}

func getSrvURL() string {
	s, f := os.LookupEnv("SERVER_ADDRESS")
	if f && s != "" {
		return s
	}

	flag.StringVar(&s, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	flag.Parse()

	return s
}

func validateOptions(options *Options) {
	if !strings.HasSuffix(options.ShortURL, "/") {
		options.ShortURL += "/"
	}

	if !strings.HasPrefix(options.ShortURL, "http://") {
		options.ShortURL = "http://" + options.ShortURL
	}
}
