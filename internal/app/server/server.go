package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/mogu10/shortlink/internal/app/controllers"
	"log"
	"net/http"
)

type Server struct {
	serverAddress string
	app           *controllers.App
}

func (s *Server) Run() {
	router := chi.NewRouter()

	router.Post("/", s.app.HandlerPost)
	router.Get("/{id}", s.app.HandlerGet)

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
