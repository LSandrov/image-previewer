package handler

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type FillRequest struct {
	width  int
	height int
	url    string
}

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	request := &FillRequest{}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	err := h.validateFillHandler(vars, request)
	if err != nil {
		w.WriteHeader(502)
		return
	}

	img, err := h.svc.Fill(ctx, request.width, request.height, request.url)
	if err != nil {
		w.WriteHeader(502)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(img)))
	if _, err := w.Write(img); err != nil {
		w.WriteHeader(502)
		h.l.Error().Msg("Невозможно обработать изображение")
	}
}

func (h *Handlers) validateFillHandler(vars map[string]string, r *FillRequest) (err error) {
	if r.width, err = strconv.Atoi(vars["width"]); err != nil {
		return errors.New("поле width должно быть целочисленным")
	}
	if r.height, err = strconv.Atoi(vars["height"]); err != nil {
		return errors.New("поле width должно быть целочисленным")
	}

	imageUrl, err := url.ParseRequestURI(vars["imageUrl"])
	if err != nil {
		return errors.New("поле imageUrl должно быть ссылкой")
	}

	r.url = imageUrl.String()

	return nil
}
