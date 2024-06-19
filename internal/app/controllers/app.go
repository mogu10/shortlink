package controllers

type App struct {
	shortAddress string
}

func New(shAd string) *App {
	return &App{
		shortAddress: shAd,
	}
}
