package handler

import (
	"image-previewer/pkg/previewer"

	"github.com/rs/zerolog"
)

type Handlers struct {
	l   zerolog.Logger
	svc previewer.Service
}

func NewHandlers(l zerolog.Logger, svc previewer.Service) *Handlers {
	return &Handlers{l: l, svc: svc}
}
