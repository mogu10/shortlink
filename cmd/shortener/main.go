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
		if options.StoragePath == "" {
			return storage.InitDefaultStorage()
		}
		return storage.InitFileStorage(options.StoragePath)
	}()

	if err != nil {
		log.Fatal(err.Error())
	}

	application := controllers.NewApp(
		controllers.WithShortAddress(options.ShortURL),
		controllers.WithStorage(st),
		controllers.WithDatabaseConnection(options.DataBaseConnection))

	serv := server.New(options.ServerURL, application)

	// запуск сервера
	serv.Run()
}
