package e2e

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		Url    string
		Status int
	}{
		// Обычный успешный кейс
		{
			Url:    "/fill/200/200/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg",
			Status: http.StatusOK,
		},
		// удаленный сервер не существует
		{
			Url:    "/fill/200/200/raw.123.jpg",
			Status: http.StatusBadGateway,
		},
		// удаленный сервер существует, но изображение не найдено (404 Not Found)
		{
			Url:    "/fill/200/200/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/blabla.jpg",
			Status: http.StatusBadGateway,
		},
		// Размер на выходе больше, чем исходное значение (масштабирование)
		{
			Url:    "/fill/2000/2000/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg",
			Status: http.StatusOK,
		},
		// Не валидные параметры обрезки
		{
			Url:    "/fill/width/height/raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/_gopher_original_1024x504.jpg",
			Status: http.StatusNotFound,
		},
		// удаленный сервер существует, но изображение не изображение
		{
			Url:    "/fill/200/200/raw.githubusercontent.com/LSandrov/image-previewer/master/Makefile",
			Status: http.StatusBadGateway,
		},
		// ошибка валидации
		{
			Url:    "/fill/200/200/user^:passwo^rd@foo.com/",
			Status: http.StatusBadRequest,
		},
	}

	for k, tt := range tests {
		q := tt
		t.Run(fmt.Sprintf("%s %d", q.Url, k), func(t *testing.T) {
			t.Parallel()
			resp, err := c.Get(buildUrl(q.Url))
			require.NoError(t, err)
			require.Equal(t, q.Status, resp.StatusCode)
			_, err = readResponse(resp)
			require.NoError(t, err)
		})
	}
}
