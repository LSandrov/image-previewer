package handler

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type FillRequest struct {
	width  int
	height int
	url    string
}

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
	defer cancel()

	vars := mux.Vars(r)

	request, err := h.parseFillHandlerVars(vars)
	if err != nil {
		w.WriteHeader(502)
		return
	}

	fillResponse, err := h.svc.Fill(ctx, request.width, request.height, request.url)
	if err != nil {
		w.WriteHeader(502)
		return
	}

	for name, values := range fillResponse.Headers {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(fillResponse.Img)))
	if _, err := w.Write(fillResponse.Img); err != nil {
		w.WriteHeader(502)
		h.l.Error().Msg("Невозможно обработать изображение")
	}
}

func (h *Handlers) parseFillHandlerVars(vars map[string]string) (r *FillRequest, err error) {
	if r.width, err = strconv.Atoi(vars["width"]); err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}
	if r.height, err = strconv.Atoi(vars["height"]); err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	imageUrl, err := url.ParseRequestURI(vars["imageUrl"])
	if err != nil {
		return nil, errors.New("поле imageUrl должно быть ссылкой")
	}

	r.url = imageUrl.String()

	return r, nil
}
