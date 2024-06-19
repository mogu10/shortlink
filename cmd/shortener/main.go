package main

import (
	"github.com/mogu10/shortlink/internal/app/config"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"github.com/mogu10/shortlink/internal/app/server"
)

func main() {
	options := config.Get()

	application := controllers.New(options.ShortURL)
	serv := server.New(options.ServerURL, application)

	// запуск сервера
	serv.Run()
}
