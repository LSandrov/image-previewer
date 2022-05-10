package internal

import (
	"net/http"
	"time"

	"github.com/LSandrov/image-previewer/internal/handler"
	"github.com/LSandrov/image-previewer/pkg/cache"
	"github.com/LSandrov/image-previewer/pkg/previewer"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type App struct {
	l             zerolog.Logger
	cacheCapacity int
}

func NewApp(l zerolog.Logger, cacheCapacity int) *App {
	return &App{
		l:             l,
		cacheCapacity: cacheCapacity,
	}
}

func (a *App) Run() {
	r := mux.NewRouter()
	resizedCache := cache.NewCache(a.cacheCapacity)
	downloadedCache := cache.NewCache(a.cacheCapacity)
	downloader := previewer.NewDefaultImageDownloader()
	svc := previewer.NewDefaultService(a.l, downloader, resizedCache, downloadedCache)
	handlers := handler.NewHandlers(a.l, svc)

	r.HandleFunc("/fill/{width:[0-9]+}/{height:[0-9]+}/{imageURL:.*}", handlers.FillHandler)
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	if err := srv.ListenAndServe(); err != nil {
		a.l.Fatal().Err(err).Msg("error starting http server")
	}
}
