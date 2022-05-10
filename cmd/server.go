package main

import (
	"image-previewer/internal"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var shaCommit = "local"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l := log.With().Str("sha_commit", shaCommit).Logger()

	if err := godotenv.Load(); err != nil {
		l.Fatal().Err(err).Msg("Error loading .env file")
	}

	cacheCapacity, err := strconv.Atoi(os.Getenv("LRU_CACHE_CAPACITY"))
	if err != nil {
		l.Fatal().Err(err).Msg("Error getting LRU_CACHE_CAPACITY from .env file")
	}

	app := internal.NewApp(l, cacheCapacity)
	app.Run()
}
