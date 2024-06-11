package controllers

type App struct {
	ShortAddress string
}

func New(shAd string) *App {
	return &App{
		ShortAddress: shAd,
	}
}
