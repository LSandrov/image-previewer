package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"image-previewer/pkg/previewer"
	mock_previewer "image-previewer/pkg/previewer/mock"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestHandlers_FillHandler_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_previewer.NewMockService(ctrl)
	l := log.With().Logger()

	image1 := loadImage("gopher_200x700.jpg")
	image2 := loadImage("gopher_1024x252.jpg")

	tests := []struct {
		name         string
		width        int64
		height       int64
		url          string
		response     string
		fillResponse *previewer.FillResponse
		err          error
		httpStatus   int64
	}{
		{
			name:         "good response",
			width:        200,
			height:       300,
			url:          "http://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_200x700.jpg",
			response:     string(image1),
			fillResponse: &previewer.FillResponse{Img: image1},
			httpStatus:   http.StatusOK,
		},
		{
			name:         "good response 2",
			width:        400,
			height:       500,
			url:          "http://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_1024x252.jpg",
			response:     string(image2),
			fillResponse: &previewer.FillResponse{Img: image2},
			httpStatus:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req = mux.SetURLVars(req, map[string]string{
				"width":    strconv.Itoa(int(tt.width)),
				"height":   strconv.Itoa(int(tt.height)),
				"imageURL": tt.url,
			})

			fillParams := previewer.NewFillParams(req.Context(), int(tt.width), int(tt.height), tt.url, req.Header)
			mockService.EXPECT().Fill(fillParams).Return(tt.fillResponse, tt.err)
			h := &Handlers{
				l:   l,
				svc: mockService,
			}

			w := httptest.NewRecorder()
			h.FillHandler(w, req)
			require.Equal(t, int(tt.httpStatus), w.Result().StatusCode)
			require.Equal(t, strings.TrimSpace(w.Body.String()), tt.response)
		})
	}
}

func TestHandlers_FillHandler_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mock_previewer.NewMockService(ctrl)
	l := log.With().Logger()

	tests := []struct {
		name         string
		width        int64
		height       int64
		url          string
		response     string
		fillResponse *previewer.FillResponse
		err          error
		httpStatus   int64
	}{
		{
			name:       "validation error",
			width:      300,
			height:     400,
			url:        "http://user^:passwo^rd@foo.com/",
			response:   "Ошибка при валидации входных данных",
			httpStatus: http.StatusBadRequest,
		},
		{
			name:         "fill error",
			width:        300,
			height:       400,
			url:          "http://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_200x700.jpg",
			response:     "Невозможно обработать изображение",
			fillResponse: nil,
			httpStatus:   http.StatusBadGateway,
			err:          errors.New("ошибка"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req = mux.SetURLVars(req, map[string]string{
				"width":    strconv.Itoa(int(tt.width)),
				"height":   strconv.Itoa(int(tt.height)),
				"imageURL": tt.url,
			})

			if tt.fillResponse != nil || tt.err != nil {
				fillParams := previewer.NewFillParams(req.Context(), int(tt.width), int(tt.height), tt.url, req.Header)
				mockService.EXPECT().Fill(fillParams).Return(tt.fillResponse, tt.err)
			}

			h := &Handlers{
				l:   l,
				svc: mockService,
			}

			w := httptest.NewRecorder()

			h.FillHandler(w, req)
			require.Equal(t, int(tt.httpStatus), w.Result().StatusCode)
			require.Equal(t, strings.TrimSpace(w.Body.String()), tt.response)
		})
	}
}
