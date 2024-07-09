package config

import (
	"flag"
	"os"
	"strings"
)

type ServiceOptions struct {
	ServerURL   string
	ShortURL    string
	StoragePath string
}

func Get() *ServiceOptions {
	options := &ServiceOptions{
		ServerURL:   getEnv("SERVER_ADDRESS"),
		ShortURL:    getEnv("BASE_URL"),
		StoragePath: getEnv("FILE_STORAGE_PATH"),
	}

	serv := ""
	short := ""
	stgePath := ""
	flag.StringVar(&serv, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	flag.StringVar(&short, "b", "http://localhost:8080/", "базовый адрес результирующего шортлинка")
	flag.StringVar(&stgePath, "f", "/tmp/short-url-db.json", "путь до файла/хранилища")
	flag.Parse()

	if options.ShortURL == "" {
		options.ShortURL = short
	}

	if options.ServerURL == "" {
		options.ServerURL = serv
	}

	if options.StoragePath == "" {
		options.StoragePath = stgePath
	}

	validateOptions(options)

	return options
}

func getEnv(key string) string {
	if value, found := os.LookupEnv(key); found && value != "" {
		return value
	}
	return ""
}

func validateOptions(options *ServiceOptions) {
	if !strings.HasSuffix(options.ShortURL, "/") {
		options.ShortURL += "/"
	}

	if !strings.HasPrefix(options.ShortURL, "http://") {
		options.ShortURL = "http://" + options.ShortURL
	}
}
