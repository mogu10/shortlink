package config

import (
	"flag"
	"strings"
)

type Options struct {
	ServerURL string
	ShortURL  string
}

func ParseArgs() *Options {
	options := new(Options)

	if flag.Lookup("a") == nil {
		flag.StringVar(&options.ServerURL, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	}

	if flag.Lookup("b") == nil {
		flag.StringVar(&options.ShortURL, "b", "http://localhost:8080/", "базовый адрес результирующего шортлинка")
	}

	flag.Parse()

	validateOptions(options)

	return options
}

func validateOptions(options *Options) {
	if !strings.HasSuffix(options.ShortURL, "/") {
		options.ShortURL += "/"
	}

	if !strings.HasPrefix(options.ShortURL, "http://") {
		options.ShortURL = "http://" + options.ShortURL
	}
}
