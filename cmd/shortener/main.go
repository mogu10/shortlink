package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strings"
)

var links = make(map[string]string)

const host string = "http://localhost:8080/"

func main() {

	router := chi.NewRouter()

	router.Post("/", postLink)
	router.Get("/{id}", getLink)

	http.ListenAndServe(":8080", router)
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

	link := host + (string(short))

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Add("Content-Type", "text/plain")
	writer.Write([]byte(link))
}

func getLink(writer http.ResponseWriter, request *http.Request) {

	//провяем, что метод GET
	if request.Method != http.MethodGet {
		http.Error(writer, "Only GET allowed", http.StatusBadRequest)
		return
	}

	//вытаскиваем path из урла
	path := request.URL.Path
	path = strings.ReplaceAll(path, "/", "")

	if len(path) == 0 {
		http.Error(writer, "Path is empty", http.StatusBadRequest)
		return
	}

	link, err := findShortLink([]byte(path))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(writer, request, string(link), http.StatusTemporaryRedirect)

}

func createShortLink(body []byte) ([]byte, error) {

	md5Hash := md5.Sum(body)
	hash := hex.EncodeToString(md5Hash[:])
	shortHash := hash[0:8]

	err := saveLink(shortHash, body)
	if err != nil {
		return nil, err
	}

	return []byte(links[string(body)]), nil
}

func saveLink(hash string, body []byte) error {
	links[string(body)] = hash

	return nil
}

func findShortLink(path []byte) ([]byte, error) {

	for key, value := range links {
		if value == string(path) {
			return []byte(key), nil
		}
	}

	return nil, errors.New("invalid path")
}
