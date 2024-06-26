package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"github.com/mogu10/shortlink/internal/logger"
	"log"
	"net/http"
)

type Server struct {
	serverAddress string
	app           *controllers.App
}

func (s *Server) Run() {
	logger.Initialize("debug")

	router := chi.NewRouter()

	router.Post("/", logger.RequestLogger(s.app.HandlerPost))
	router.Post("/api/shorten", logger.RequestLogger(s.app.HandlerPostJson))

	router.Get("/{id}", logger.RequestLogger(s.app.HandlerGet))

	err := http.ListenAndServe(s.serverAddress, router)

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
