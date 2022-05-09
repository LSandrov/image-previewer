package previewer

import (
	"context"
	"github.com/rs/zerolog"
	"image-previewer/pkg/cache"
	"reflect"
	"testing"
)

func TestDefaultService_Fill(t *testing.T) {
	type fields struct {
		l          zerolog.Logger
		downloader ImageDownloader
		cache      cache.Cache
	}
	type args struct {
		ctx    context.Context
		width  int
		height int
		imgURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *FillResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &DefaultService{
				l:          tt.fields.l,
				downloader: tt.fields.downloader,
				cache:      tt.fields.cache,
			}
			got, err := svc.Fill(tt.args.ctx, tt.args.width, tt.args.height, tt.args.imgURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fill() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fill() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultService_resize(t *testing.T) {
	type fields struct {
		l          zerolog.Logger
		downloader ImageDownloader
		cache      cache.Cache
	}
	type args struct {
		img    []byte
		width  int
		height int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &DefaultService{
				l:          tt.fields.l,
				downloader: tt.fields.downloader,
				cache:      tt.fields.cache,
			}
			got, err := svc.resize(tt.args.img, tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("resize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultService(t *testing.T) {
	type args struct {
		l          zerolog.Logger
		downloader ImageDownloader
		c          cache.Cache
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultService(tt.args.l, tt.args.downloader, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultService() = %v, want %v", got, tt.want)
			}
		})
	}
}
