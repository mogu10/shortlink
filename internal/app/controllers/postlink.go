package controllers

import (
	"encoding/json"
	"errors"
	"github.com/mogu10/shortlink/internal/app/service"
	"io"
	"log"
	"net/http"

	"github.com/mogu10/shortlink/internal/app/storage"
)

func (a *App) HandlerPost(writer http.ResponseWriter, request *http.Request) {
	body, err := getBodyFromRequest(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
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

func (a *App) HandlerPostJson(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	defer request.Body.Close()

	requestJson := RequestFields{Url: ""}

	decoder.Decode(&requestJson.Url)
	short, err := createShortLink([]byte(requestJson.Url))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	link := a.shortAddress + (string(short))
	responseJson := ResponseFields{Result: link}
	resp, err := json.Marshal(responseJson)

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.Write(resp)
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

func getBodyFromRequest(request *http.Request) ([]byte, error) {
	if request.Method != http.MethodPost {
		return nil, errors.New("only POST allowed")
	}

	body, err := io.ReadAll(request.Body)
	request.Body.Close()

	if err != nil {
		return nil, errors.New("something wrong with body")
	}

	return body, nil
}
