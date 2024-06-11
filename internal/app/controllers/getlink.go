package controllers

import (
	"github.com/mogu10/shortlink/internal/app/storage"
	"net/http"
	"strings"
)

func (a *App) HandlerGet(writer http.ResponseWriter, request *http.Request) {

	// провяем, что метод GET
	if request.Method != http.MethodGet {
		http.Error(writer, "Only GET allowed", http.StatusBadRequest)
		return
	}

	// вытаскиваем path из урла
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

func findShortLink(path []byte) ([]byte, error) {

	link, err := storage.LoadLink(path)
	if err != nil {
		return nil, err
	}

	return link, nil
}
