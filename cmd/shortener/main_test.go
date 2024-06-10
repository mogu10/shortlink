package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostLink(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name  string
		want  want
		url   string
		short string
	}{
		{
			name:  "positive test #1",
			url:   "https://yandex.ru",
			short: "http://localhost:8080/e9db20b2",
			want: want{
				code:        201,
				response:    `{"status":"ok"}`,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// делаем тестовый POST запрос
			w, err := createPostLinkRequest(test.url)
			if err != nil {
				t.Fatal(err)
			}

			res := w.Result()

			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			assert.Equal(t, test.short, string(resBody))

			// получаем и проверяем заголовок
			header := w.Header()
			require.NoError(t, err)
			assert.Equal(t, test.want.contentType, header.Get("Content-Type"))
		})
	}
}

func TestGetLink(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}
	tests := []struct {
		name  string
		want  want
		url   string
		short string
	}{
		{
			name:  "positive test #1",
			url:   "https://yandex.ru",
			short: "e9db20b2",
			want: want{
				code:        307,
				contentType: "text/html; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// делаем тестовый POST запрос
			_, err := createPostLinkRequest(test.url)

			if err != nil {
				t.Fatal(err)
			}

			request := httptest.NewRequest(http.MethodGet, "/"+test.short, nil)
			request.Header.Set("Content-Type", "text/plain")

			// создаём новый Recorder
			w := httptest.NewRecorder()
			getLink(w, request)
			res := w.Result()

			// проверяем код ответа
			assert.Equal(t, test.want.code, res.StatusCode)

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			//проверяем, что в теле содержится исходная ссылка
			assert.Contains(t, string(resBody), test.url)

			header := w.Header()

			require.NoError(t, err)
			assert.Equal(t, test.want.contentType, header.Get("Content-Type"))
		})
	}
}

func createPostLinkRequest(url string) (*httptest.ResponseRecorder, error) {
	body := strings.NewReader(url)
	request := httptest.NewRequest(http.MethodPost, "/", body)
	request.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()
	postLink(w, request)

	return w, nil
}
