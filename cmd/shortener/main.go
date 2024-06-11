package main

import (
	"github.com/mogu10/shortlink/internal/app/config"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"github.com/mogu10/shortlink/internal/app/server"
)

func main() {
	options := config.Get()
	//options := config.ParseArgs()

	a := controllers.New(options.ShortURL)
	s := server.New(options.ServerURL, a)

	// запуск сервера
	s.Run()
}
