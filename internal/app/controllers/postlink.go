package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/mogu10/shortlink/internal/app/service"
	"github.com/mogu10/shortlink/internal/app/storage"
	"io"
	"log"
	"net/http"
)

func (a *App) HandlerPost(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Only POST allowed", http.StatusBadRequest)
		return
	}

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

func (a *App) HandlerPostJSON(writer http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer
	var requestFiels RequestFields

	_, err := buf.ReadFrom(request.Body)
	defer request.Body.Close()

	if err != nil {
		http.Error(writer, "Something wrong with body", http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &requestFiels); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	short, err := createShortLink([]byte(requestFiels.URL))

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	link := a.shortAddress + (string(short))
	responseJSON := ResponseFields{Result: link}
	response, err := json.Marshal(responseJSON)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.Write(response)
}

func createShortLink(text []byte) ([]byte, error) {
	shortHash := service.HashText(text)
	err := storage.SaveLink(shortHash, text)

	if err != nil {
		return nil, err
	}

	log.Println("сохранена пара: " + shortHash + " - " + string(text))

	return []byte(shortHash), nil
}
