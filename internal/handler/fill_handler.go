package handler

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"image-previewer/pkg/previewer"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	fillParams, err := h.parseFillHandlerVars(r.Context(), mux.Vars(r), r.Header)
	if err != nil {
		w.WriteHeader(500)
		h.l.Err(err).Msg("Ошибка при валидации входных данных")
		w.Write([]byte("Ошибка при валидации входных данных"))
		return
	}

	fillResponse, err := h.svc.Fill(fillParams)
	if err != nil {
		w.WriteHeader(500)
		h.l.Err(err).Msg("Невозможно обработать изображение")
		w.Write([]byte("Невозможно обработать изображение"))
		return
	}

	for name, values := range fillResponse.Headers {
		for _, value := range values {
			w.Header().Set(name, value)
		}
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(fillResponse.Img)))
	if _, err := w.Write(fillResponse.Img); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Проблемы с обработкой ответа"))
		h.l.Err(err).Msg("Проблемы с обработкой ответа")
		return
	}
}

func (h *Handlers) parseFillHandlerVars(ctx context.Context, vars map[string]string, headers map[string][]string) (*previewer.FillParams, error) {
	width, err := strconv.Atoi(vars["width"])
	if err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	height, err := strconv.Atoi(vars["height"])
	if err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	imageUrl, err := url.Parse(vars["imageUrl"])
	if err != nil {
		return nil, errors.New("поле imageUrl должно быть ссылкой")
	}

	//Условие тз: Работаем только с HTTP.
	imageUrl.Scheme = "http"

	return previewer.NewFillParams(ctx, width, height, imageUrl.String(), headers), nil
}
