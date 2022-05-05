package handler

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"image-previewer/pkg/previewer"
	mock_previewer "image-previewer/pkg/previewer/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const ImageURL = "https://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/"

func TestHandlers_FillHandler(t *testing.T) {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()
	mockService := mock_previewer.NewMockService(ctrl)
	l := log.With().Logger()

	image1 := loadImage("gopher_200x700.jpg")
	image2 := loadImage("gopher_1024x252.jpg")

	tests := []struct {
		name         string
		width        string
		height       string
		url          string
		response     string
		fillResponse *previewer.FillResponse
	}{
		{
			name:         "good",
			width:        "200",
			height:       "300",
			url:          "https://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_200x700.jpg",
			response:     string(image1),
			fillResponse: &previewer.FillResponse{Img: image1},
		},
		{
			name:         "good",
			width:        "200",
			height:       "300",
			url:          "https://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_1024x252.jpg",
			response:     string(image2),
			fillResponse: &previewer.FillResponse{Img: image2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req = mux.SetURLVars(req, map[string]string{
				"width":    tt.width,
				"height":   tt.height,
				"imageUrl": tt.url,
			})

			mockService.EXPECT().Fill(req.Context(), tt.width, tt.height, tt.url).Return(tt.fillResponse, nil)
			h := &Handlers{
				l:   l,
				svc: mockService,
			}

			w := httptest.NewRecorder()

			h.FillHandler(w, req)
			require.Equal(t, http.StatusOK, w.Result().StatusCode)
			require.Equal(t, strings.TrimSpace(w.Body.String()), tt.response)
			require.Equal(t, w.Header().Get("Content-Type"), "image/jpeg")
		})
	}
}
