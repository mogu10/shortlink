package config

import (
	"flag"
	"fmt"
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

	if options.ShortURL == "" {
		flag.StringVar(
			&options.ShortURL,
			"a",
			"localhost:8080",
			"адрес запуска HTTP-сервера")
	}

	if options.ServerURL == "" {
		flag.StringVar(
			&options.ServerURL,
			"b",
			"http://localhost:8080/",
			"базовый адрес результирующего шортлинка")

	}

	flag.Parse()
	fmt.Println(options.ServerURL)
	validateOptions(options)

	return options
}

func getShURL() string {
	s, f := os.LookupEnv("BASE_URL")
	if f && s != "" {
		return s
	}

	return ""
}

func getSrvURL() string {
	s, f := os.LookupEnv("SERVER_ADDRESS")
	if f && s != "" {
		return s
	}

	return ""
}

func validateOptions(options *Options) {
	if !strings.HasSuffix(options.ShortURL, "/") {
		options.ShortURL += "/"
	}

	if !strings.HasPrefix(options.ShortURL, "http://") {
		options.ShortURL = "http://" + options.ShortURL
	}
}
