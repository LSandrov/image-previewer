package handler

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
)

type FillRequest struct {
	width  int
	height int
	url    string
}

func (h *Handlers) FillHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	request, err := h.parseFillHandlerVars(vars)
	if err != nil {
		w.WriteHeader(500)
		h.l.Err(err).Msg("Ошибка при валидации входных данных")
		w.Write([]byte("Ошибка при валидации входных данных"))
		return
	}

	fillResponse, err := h.svc.Fill(r.Context(), request.width, request.height, request.url)
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

func (h *Handlers) parseFillHandlerVars(vars map[string]string) (*FillRequest, error) {
	r := &FillRequest{}

	width, err := strconv.Atoi(vars["width"])
	if err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	r.width = width

	height, err := strconv.Atoi(vars["height"])
	if err != nil {
		return nil, errors.New("поле width должно быть целочисленным")
	}

	r.height = height

	imageUrl, err := url.Parse(vars["imageUrl"])
	if err != nil {
		return nil, errors.New("поле imageUrl должно быть ссылкой")
	}

	imageUrl.Scheme = "https"

	r.url = imageUrl.String()

	return r, nil
}
