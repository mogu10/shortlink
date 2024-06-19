package controllers

import (
	"github.com/mogu10/shortlink/internal/app/service"
	"io"
	"net/http"

	"github.com/mogu10/shortlink/internal/app/storage"
)

func (a *App) HandlerPost(writer http.ResponseWriter, request *http.Request) {
	// провяем, что метод POST
	if request.Method != http.MethodPost {
		http.Error(writer, "Only POST allowed", http.StatusBadRequest)
		return
	}

	// вытаскиваем body из реквеста
	body, err := io.ReadAll(request.Body)
	request.Body.Close()

	if err != nil {
		http.Error(writer, "Something wrong with body", http.StatusBadRequest)
		return
	}

	short, err := createShortLink(body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	link := a.shortAddress + (string(short))

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "text/plain")
	writer.Write([]byte(link))
}

func createShortLink(body []byte) ([]byte, error) {
	shortHash := service.HashText(body)
	err := storage.SaveLink(shortHash, body)

	if err != nil {
		return nil, err
	}

	return []byte(shortHash), nil
}
