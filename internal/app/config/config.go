package config

import "flag"

type Options struct {
	ServerUrl string
	ShortUrl  string
}

func ParseArgs() *Options {
	options := new(Options)

	flag.StringVar(&options.ServerUrl, "a", "localhost:8000", "адрес запуска HTTP-сервера")
	flag.StringVar(&options.ShortUrl, "b", "http://localhost:8000/", "базовый адрес результирующего шортлинка")
	flag.Parse()

	return options
}
