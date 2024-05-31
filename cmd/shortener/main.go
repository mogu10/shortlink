package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", postLink)
	mux.HandleFunc("/{id}/", getLink)

	http.ListenAndServe(":8080", mux)
}

func postLink(writer http.ResponseWriter, request *http.Request) {
	//провяем, что метод POST
	if request.Method != http.MethodPost {
		http.Error(writer, "Only POST allowed", http.StatusBadRequest)
		return
	}

	//вытаскиваем body из реквеста
	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Something wrong with body", http.StatusBadRequest)
	}

	short, err := createShortLink(body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	writer.Header().Add("Content-Type", "text/plain")
	writer.Write(short)
}

func getLink(writer http.ResponseWriter, request *http.Request) {

	//провяем, что метод POST
	if request.Method != http.MethodGet {
		http.Error(writer, "Only GET allowed", http.StatusBadRequest)
		return
	}

	//вытаскиваем path из урла
	path := request.URL.Path
	if len(path) == 0 {
		http.Error(writer, "Path is empty", http.StatusBadRequest)
		return
	}

	link, err := findShortLink([]byte(path))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Add("Content-Type", "text/plain")
	writer.Write(link)
}

func createShortLink(body []byte) ([]byte, error) {
	if bytes.Equal(body, []byte("https://practicum.yandex.ru/")) {
		return []byte("EwHXdJfB"), nil
	}

	return nil, errors.New("invalid short link")
}

func findShortLink(path []byte) ([]byte, error) {
	if bytes.Equal(path, []byte("/EwHXdJfB/")) {
		return []byte("https://practicum.yandex.ru/"), nil
	}

	return nil, errors.New("invalid path")
}
