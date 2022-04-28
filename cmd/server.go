package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"image-previewer/internal"
	"os"
	"strconv"
)

var shaCommit = "local"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := log.With().Str("sha_commit", shaCommit).Logger()

	err := godotenv.Load()
	if err != nil {
		l.Fatal().Err(err).Msg("Error loading .env file")
	}

	cacheCapacity, err := strconv.Atoi(os.Getenv("LRU_CACHE_CAPACITY"))
	if err != nil {
		l.Fatal().Err(err).Msg("Error getting LRU_CACHE_CAPACITY from .env file")
	}

	app := internal.NewApp(l, cacheCapacity)
	app.Run()
}
