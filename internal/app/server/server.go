package server

import (
	"github.com/mogu10/shortlink/internal/app/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	ServerAddress string
	App           *controllers.App
}

func (s *Server) Run() {
	router := chi.NewRouter()

	router.Post("/", s.App.HandlerPost)
	router.Get("/{id}", s.App.HandlerGet)

	http.ListenAndServe(s.ServerAddress, router)
}

func New(srvAd string, a *controllers.App) *Server {
	return &Server{
		ServerAddress: srvAd,
		App:           a,
	}
}
