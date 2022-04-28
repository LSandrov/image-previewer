package handler

import (
	"github.com/rs/zerolog"
	"image-previewer/pkg/previewer"
)

type Handlers struct {
	l   zerolog.Logger
	svc previewer.Service
}

func NewHandlers(l zerolog.Logger, svc previewer.Service) *Handlers {
	return &Handlers{l: l, svc: svc}
}
