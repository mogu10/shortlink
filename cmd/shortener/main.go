package main

import (
	"github.com/mogu10/shortlink/internal/app/config"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"github.com/mogu10/shortlink/internal/app/server"
	"github.com/mogu10/shortlink/internal/app/storage"
	"log"
)

func main() {
	var st storage.Storage
	var err error

	options := config.Get()

	if options.StoragePath == "" {
		st, err = storage.InitDefaultStorage()
	} else {
		st, err = storage.InitFileStorage(options.StoragePath)
	}

	if err != nil {
		log.Panicf(err.Error())
	}

	application := controllers.New(options.ShortURL, st)
	serv := server.New(options.ServerURL, application)

	// запуск сервера
	serv.Run()
}
