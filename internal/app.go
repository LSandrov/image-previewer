package internal

import (
	"image-previewer/pkg/previewer"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"image-previewer/internal/handler"
	"image-previewer/pkg/cache"
)

type App struct {
	l zerolog.Logger
	c cache.Cache
}

func NewApp(l zerolog.Logger, cacheCapacity int) *App {
	return &App{
		l: l,
		c: cache.NewCache(cacheCapacity),
	}
}

func (a *App) Run() {
	r := mux.NewRouter()
	downloader := previewer.NewDefaultImageDownloader(a.l)
	svc := previewer.NewDefaultService(a.l, downloader, a.c)
	handlers := handler.NewHandlers(a.l, svc)

	r.HandleFunc("/fill/{width:[0-9]+}/{height:[0-9]+}/{imageUrl:.*}", handlers.FillHandler)
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	srv.ListenAndServe()
}
