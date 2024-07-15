package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mogu10/shortlink/internal/app/config"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"github.com/mogu10/shortlink/internal/app/server"
	"github.com/mogu10/shortlink/internal/app/storage"
	"log"
)

func main() {
	options := config.Get()

	st, err := func() (storage.Storage, error) {
		if options.DataBaseConnection == "" {
			if options.StoragePath == "" {
				return storage.InitDefaultStorage()
			}

			return storage.InitFileStorage(options.StoragePath)
		}

		return storage.Connection(options.DataBaseConnection)
	}()

	if err != nil {
		log.Fatal(err.Error())
	}

	application := controllers.NewApp(
		controllers.WithShortAddress(options.ShortURL),
		controllers.WithStorage(st))

	serv := server.New(options.ServerURL, application)

	// запуск сервера
	serv.Run()
}
