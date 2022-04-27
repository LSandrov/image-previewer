package internal

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"image-previewer/internal/handler"
	"image-previewer/pkg/cache"
	"net/http"
)

type App struct {
	l zerolog.Logger
	c cache.Cache
}

func NewApp(l zerolog.Logger, c cache.Cache) *App {
	return &App{l: l, c: c}
}

func (a *App) Run() {
	r := mux.NewRouter()
	handlers := handler.NewHandlers()
	r.HandleFunc("/fill/{width:[0-9]}/{height:[0-9]}/{imageUrl:.*}", handlers.FillHandler)
	http.Handle("/", r)
}
