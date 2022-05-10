package previewer

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"time"
)

const ImageURL = "http://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/"

func TestDefaultImageDownloader_DownloadByUrl(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		ctx := context.Background()

		tests := []struct {
			ctx     context.Context
			imgName string
			headers map[string][]string
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
				d := &DefaultImageDownloader{}
				gotImg, err := d.DownloadByUrl(tt.ctx, ImageURL+tt.imgName, tt.headers)
				if err != nil {
					t.Errorf("DownloadByUrl() error = %v", err)
					return
				}

				wantImg := loadImage(tt.imgName)
				if !reflect.DeepEqual(gotImg.img, wantImg) {
					t.Errorf("DownloadByUrl() gotImg = %v, want %v", gotImg, wantImg)
				}
			})
		}
	})

	t.Run("negative", func(t *testing.T) {
		ctx := context.Background()
		ctxWithTimeOut, closefn := context.WithTimeout(ctx, time.Microsecond*1)
		defer closefn()

		tests := []struct {
			ctx     context.Context
			imgName string
			url     string
			headers map[string][]string
			err     error
		}{
			{
				ctx:     ctxWithTimeOut,
				imgName: "gopher_200x700.jpg",
				url:     ImageURL,
				err:     ErrTimeout,
			},
			{
				ctx:     ctxWithTimeOut,
				imgName: "gopher_1024x252.jpg",
				url:     ImageURL,
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
				d := &DefaultImageDownloader{}
				_, err := d.DownloadByUrl(tt.ctx, tt.url+tt.imgName, tt.headers)
				require.Errorf(t, err, tt.err.Error())
			})
		}
	})
}

func TestNewDefaultImageDownloader(t *testing.T) {
	tests := []struct {
		name string
		want ImageDownloader
	}{
		{
			name: "good",
			want: &DefaultImageDownloader{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultImageDownloader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultImageDownloader() = %v, want %v", got, tt.want)
			}
		})
	}
}
