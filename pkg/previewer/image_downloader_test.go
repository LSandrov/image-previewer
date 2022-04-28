package previewer

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

const ImageUrl = "https://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/"

func TestDefaultImageDownloader_DownloadByUrl_Positive(t *testing.T) {
	logger := log.With().Str("test", "test").Logger()
	ctx := context.Background()

	tests := []struct {
		ctx     context.Context
		imgName string
	}{
		{
			ctx:     ctx,
			imgName: "gopher_200x700.jpg",
		},
		{
			ctx:     ctx,
			imgName: "gopher_1024x252.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.imgName, func(t *testing.T) {
			d := &DefaultImageDownloader{
				l: logger,
			}
			gotImg, err := d.DownloadByUrl(tt.ctx, ImageUrl+tt.imgName)
			if err != nil {
				t.Errorf("DownloadByUrl() error = %v", err)
				return
			}

			wantImg := loadImage(tt.imgName)
			if !reflect.DeepEqual(gotImg, wantImg) {
				t.Errorf("DownloadByUrl() gotImg = %v, want %v", gotImg, wantImg)
			}
		})
	}
}

func TestDefaultImageDownloader_DownloadByUrl_Negative(t *testing.T) {
	logger := log.With().Str("test", "test").Logger()
	ctx := context.Background()
	ctxWithTimeOut, _ := context.WithTimeout(ctx, time.Microsecond*1)

	tests := []struct {
		ctx     context.Context
		imgName string
		url     string
		err     error
	}{
		{
			ctx:     ctxWithTimeOut,
			imgName: "gopher_200x700.jpg",
			url:     ImageUrl,
			err:     ErrTimeout,
		},
		{
			ctx:     ctxWithTimeOut,
			imgName: "gopher_1024x252.jpg",
			url:     ImageUrl,
			err:     ErrTimeout,
		},
		{
			ctx:     ctx,
			imgName: "",
			url:     "",
			err:     errors.New("unsupported protocol scheme"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.imgName, func(t *testing.T) {
			d := &DefaultImageDownloader{
				l: logger,
			}
			_, err := d.DownloadByUrl(tt.ctx, tt.url+tt.imgName)
			require.Errorf(t, err, tt.err.Error())
		})
	}
}
