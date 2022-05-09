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
	"strconv"
	"strings"
	"testing"
)

func TestHandlers_FillHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	l := log.With().Logger()

	image1 := loadImage("gopher_200x700.jpg")
	image2 := loadImage("gopher_1024x252.jpg")

	tests := []struct {
		name         string
		width        int
		height       int
		url          string
		response     string
		fillResponse *previewer.FillResponse
		err          error
	}{
		{
			name:         "good",
			width:        200,
			height:       300,
			url:          "https://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_200x700.jpg",
			response:     string(image1),
			fillResponse: &previewer.FillResponse{Img: image1},
		},
		{
			name:         "good1",
			width:        300,
			height:       400,
			url:          "https://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/gopher_1024x252.jpg",
			response:     string(image2),
			fillResponse: &previewer.FillResponse{Img: image2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
			req = mux.SetURLVars(req, map[string]string{
				"width":    strconv.Itoa(tt.width),
				"height":   strconv.Itoa(tt.height),
				"imageUrl": tt.url,
			})

			mockService := mock_previewer.NewMockService(ctrl)
			mockService.EXPECT().Fill(req.Context(), tt.width, tt.height, tt.url).Return(tt.fillResponse, tt.err)
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
