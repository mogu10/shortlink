package controllers

import "github.com/mogu10/shortlink/internal/app/storage"

type App struct {
	shortAddress    string
	storage         storage.Storage
	dbConnectionStr string
}

func NewApp(opts ...func(*App)) *App {
	app := &App{}
	for _, opt := range opts {
		opt(app)
	}
	return app
}

func WithShortAddress(shortAddress string) func(*App) {
	return func(app *App) { app.shortAddress = shortAddress }
}

func WithStorage(storage storage.Storage) func(*App) {
	return func(app *App) { app.storage = storage }
}

func WithDatabaseConnection(connectionStr string) func(*App) {
	return func(app *App) { app.dbConnectionStr = connectionStr }
}
