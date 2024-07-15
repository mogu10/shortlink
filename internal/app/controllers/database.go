package controllers

import (
	"log"
	"net/http"
)

func (a *App) PingDB(writer http.ResponseWriter, request *http.Request) {
	result, err := a.storage.ConnectionCheck()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("База не пингуется значение %v", result, err)
	}

	if result {
		writer.WriteHeader(http.StatusOK)
	}
}
