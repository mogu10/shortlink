package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"github.com/mogu10/shortlink/internal/compression"
	"github.com/mogu10/shortlink/internal/logger"
	"log"
	"net/http"
)

type Server struct {
	serverAddress string
	app           *controllers.App
}

func (s *Server) Run() {
	// инитим логгер
	l, err := logger.Initialize("debug")
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Post("/", s.app.HandlerPost)
	router.Post("/api/shorten", s.app.HandlerPostJSON)

	router.Get("/{id}", s.app.HandlerGet)

	router.Get("/ping", s.app.PingDB)

	err = http.ListenAndServe(
		s.serverAddress,
		l.RequestLoggerMV(
			compression.GzipMV(
				router),
		),
	)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func New(srvAd string, a *controllers.App) *Server {
	return &Server{
		serverAddress: srvAd,
		app:           a,
	}
}
