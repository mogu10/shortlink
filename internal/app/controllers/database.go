package controllers

import (
	"net/http"
)

func (a *App) PingDB(writer http.ResponseWriter, request *http.Request) {
	result, err := a.storage.ConnectionCheck()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	if result {
		writer.WriteHeader(http.StatusOK)
	}
}
