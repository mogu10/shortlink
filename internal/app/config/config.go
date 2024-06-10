package config

import "flag"

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

	return options
}
