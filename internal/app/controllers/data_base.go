package controllers

import (
	"github.com/mogu10/shortlink/internal/app/storage"
	"log"
	"net/http"
)

func (a *App) PingDB(writer http.ResponseWriter, request *http.Request) {
	result, err := storage.ConnectionCheck(a.dbConnectionStr)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	if result {
		writer.WriteHeader(http.StatusOK)
	}
}
