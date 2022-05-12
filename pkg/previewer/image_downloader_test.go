package previewer

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const ImageURL = "http://raw.githubusercontent.com/OtusGolang/final_project/master/examples/image-previewer/"

func TestDefaultImageDownloader_DownloadByUrl_Positive(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		ctx     context.Context
		imgName string
		headers map[string][]string
	}{
		{
			ctx:     ctx,
			imgName: "_gopher_original_1024x504.jpg",
		},
		{
			ctx:     ctx,
			imgName: "gopher_1024x252.jpg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.imgName, func(t *testing.T) {
			d := &DefaultImageDownloader{}
			gotImg, err := d.DownloadByURL(tt.ctx, ImageURL+tt.imgName, tt.headers)
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
}

func TestDefaultImageDownloader_DownloadByUrl_Negative(t *testing.T) {
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
			imgName: "_gopher_original_1024x504.jpg",
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
			imgName: "Makefile",
			url:     "https://raw.githubusercontent.com/LSandrov/image-previewer/master/",
			err:     ErrImgValidate,
		},
		{
			ctx:     ctx,
			imgName: "",
			url:     string(rune(0x7f)),
			err:     errors.New("net/url: invalid control character in URL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.imgName, func(t *testing.T) {
			d := &DefaultImageDownloader{}
			_, err := d.DownloadByURL(tt.ctx, tt.url+tt.imgName, tt.headers)
			require.Errorf(t, err, tt.err.Error())
		})
	}
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

func TestDefaultImageDownloader_validate(t *testing.T) {
	type args struct {
		img []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				img: loadImage("gopher_100x100.jpg"),
			},
			wantErr: false,
		},
		{
			name: "bad format",
			args: args{
				img: loadImage("img.txt"),
			},
			wantErr: true,
		},
		{
			name: "empty img",
			args: args{
				img: loadImage("empty_img.jpg"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DefaultImageDownloader{}
			if err := d.validate(tt.args.img); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
