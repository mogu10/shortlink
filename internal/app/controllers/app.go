package controllers

type App struct {
	shortAddress string
}

type RequestFields struct {
	Url string `json:"url"`
}

type ResponseFields struct {
	Result string `json:"result"`
}

func New(shAd string) *App {
	return &App{
		shortAddress: shAd,
	}
}
