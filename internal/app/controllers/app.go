package controllers

import "github.com/mogu10/shortlink/internal/app/storage"

type App struct {
	shortAddress string
	storage      storage.Storage
}

type RequestFields struct {
	URL string `json:"url"`
}

type ResponseFields struct {
	Result string `json:"result"`
}

func New(shAd string, stge storage.Storage) *App {
	return &App{
		shortAddress: shAd,
		storage:      stge,
	}
}
