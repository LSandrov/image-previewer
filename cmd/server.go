package main

import (
	"image-previewer/internal"
	"image-previewer/pkg/cache"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var shaCommit = "local"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := log.With().Str("sha_commit", shaCommit).Logger()

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		l.Fatal().Err(err).Msg("Ошибка загрузки env файла %s")
	}

	lruCacheCapacity, ok := viper.Get("LRU_CACHE_CAPACITY").(int)
	if !ok {
		l.Fatal().Err(err).Msg("Неизвестная переменная окружения")
	}
	lruCache := cache.NewCache(lruCacheCapacity)

	app := internal.NewApp(l, lruCache)
	app.Run()
}
