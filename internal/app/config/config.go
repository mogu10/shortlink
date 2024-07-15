package config

import (
	"flag"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strings"
)

type ServiceOptions struct {
	ServerURL          string
	ShortURL           string
	StoragePath        string
	DataBaseConnection string
}

func Get() *ServiceOptions {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	options := &ServiceOptions{
		ServerURL:          getEnv("SERVER_ADDRESS"),
		ShortURL:           getEnv("BASE_URL"),
		StoragePath:        getEnv("FILE_STORAGE_PATH"),
		DataBaseConnection: getEnv("DATABASE_DSN"),
	}

	serv := ""
	short := ""
	stgePath := ""
	dataBase := ""
	flag.StringVar(&serv, "a", "", "адрес запуска HTTP-сервера")
	flag.StringVar(&short, "b", "", "базовый адрес результирующего шортлинка")
	flag.StringVar(&stgePath, "f", "", "путь до файла/хранилища")
	flag.StringVar(&dataBase, "d", "", "строка с адресом подключения к БД")
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

	if options.DataBaseConnection == "" {
		options.DataBaseConnection = dataBase
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
