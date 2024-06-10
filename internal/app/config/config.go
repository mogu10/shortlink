package config

import (
	"flag"
	"strings"
)

type Options struct {
	ServerUrl string
	ShortUrl  string
}

func ParseArgs() *Options {
	options := new(Options)

	if flag.Lookup("a") == nil {
		flag.StringVar(&options.ServerUrl, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	}

	if flag.Lookup("b") == nil {
		flag.StringVar(&options.ShortUrl, "b", "http://localhost:8080/", "базовый адрес результирующего шортлинка")
	}

	flag.Parse()

	validateOptions(options)

	return options
}

func validateOptions(options *Options) {
	if !strings.HasSuffix(options.ShortUrl, "/") {
		options.ShortUrl += "/"
	}

	if !strings.HasPrefix(options.ShortUrl, "http://") {
		options.ShortUrl = "http://" + options.ShortUrl
	}
}
