package previewer

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

var (
	ErrTimeout   = fmt.Errorf("download timeout")
	ErrReadImage = fmt.Errorf("error read image")
	ErrRequest   = fmt.Errorf("error to create request")
)

type ImageDownloader interface {
	DownloadByUrl(ctx context.Context, url string) (img []byte, err error)
}

type DefaultImageDownloader struct {
	l zerolog.Logger
}

func NewDefaultImageDownloader(l zerolog.Logger) ImageDownloader {
	return &DefaultImageDownloader{l: l}
}

func (d *DefaultImageDownloader) DownloadByUrl(ctx context.Context, url string) (img []byte, err error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		d.l.Error().Msg(ErrRequest.Error())
		return nil, err
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		if networkErr, ok := err.(net.Error); ok && networkErr.Timeout() {
			d.l.Error().Msg(ErrTimeout.Error())
			return nil, ErrTimeout
		}

		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.l.Error().Msg(ErrReadImage.Error())
		}
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}
