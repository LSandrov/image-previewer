package handler

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"image-previewer/pkg/previewer"

	"github.com/gorilla/mux"
)

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	fillParams, err := h.parseFillHandlerVars(r.Context(), mux.Vars(r), r.Header)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка при валидации входных данных"))
		return
	}

	fillResponse, err := h.svc.Fill(fillParams)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("Невозможно обработать изображение"))
		return
	}

	for name, values := range fillResponse.Headers {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}

	if _, err := w.Write(fillResponse.Img); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.l.Err(err).Msg("Проблемы с обработкой ответа")
		return
	}
}

func (h *Handlers) parseFillHandlerVars(
	ctx context.Context,
	vars map[string]string,
	headers map[string][]string,
) (*previewer.FillParams, error) {
	width, err := strconv.Atoi(vars["width"])
	if err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	height, err := strconv.Atoi(vars["height"])
	if err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	imageURL, err := url.Parse(vars["imageURL"])
	if err != nil {
		return nil, errors.New("поле imageURL должно быть ссылкой")
	}

	// Условие тз: Работаем только с HTTP.
	imageURL.Scheme = "http"

	return previewer.NewFillParams(ctx, width, height, imageURL.String(), headers), nil
}
