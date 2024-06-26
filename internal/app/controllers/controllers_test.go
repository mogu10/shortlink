package controllers

import (
	"encoding/json"
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
		route string
	}{
		{
			name:  "positive test #1",
			url:   "https://yandex.ru",
			route: "http://localhost:8080/",
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

			body := strings.NewReader(test.url)
			w, err := createPostLinkRequest(body, test.route)
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

func TestPostLinkJson(t *testing.T) {
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
		route string
	}{
		{
			name:  "positive test #1",
			url:   "https://yandex.ru",
			route: "http://localhost:8080/",
			short: "{\"result\":\"http://localhost:8080/d41d8cd9\"}",
			want: want{
				code:        201,
				response:    `{"status":"ok"}`,
				contentType: "application/json; charset=utf-8",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// делаем тестовый POST запрос

			type jsonBody struct {
				Url string `json:"url"`
			}

			jsBody, _ := json.Marshal(jsonBody{Url: test.url})

			body := strings.NewReader(string(jsBody))
			w, err := createPostLinkRequestJson(body, test.route)

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
		route string
	}{
		{
			name:  "positive test #1",
			url:   "https://yandex.ru",
			route: "http://localhost:8080/",
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
			body := strings.NewReader(test.url)
			_, err := createPostLinkRequest(body, test.route)

			if err != nil {
				t.Fatal(err)
			}

			request := httptest.NewRequest(http.MethodGet, "/"+test.short, nil)
			request.Header.Set("Content-Type", "text/plain")

			// создаём новый Recorder
			w := httptest.NewRecorder()

			app := New("http://localhost:8080/")

			app.HandlerGet(w, request)

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

func createPostLinkRequest(body *strings.Reader, route string) (*httptest.ResponseRecorder, error) {
	request := httptest.NewRequest(http.MethodPost, "/", body)
	request.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()

	app := New(route)
	app.HandlerPost(w, request)

	return w, nil
}

func createPostLinkRequestJson(body *strings.Reader, route string) (*httptest.ResponseRecorder, error) {
	request := httptest.NewRequest(http.MethodPost, "/api/shorten", body)
	request.Header.Set("Content-Type", "text/plain")

	w := httptest.NewRecorder()

	app := New(route)
	app.HandlerPostJson(w, request)

	return w, nil
}
