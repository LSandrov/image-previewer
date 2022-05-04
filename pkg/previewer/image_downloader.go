package previewer

import (
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var (
	ErrTimeout = errors.New("download timeout")
	ErrRequest = errors.New("error to create request")
)

type ImageDownloader interface {
	DownloadByUrl(ctx context.Context, url string) (*DownloadedImage, error)
}

type DefaultImageDownloader struct {
}

type DownloadedImage struct {
	img     []byte
	headers map[string][]string
}

func NewDefaultImageDownloader() ImageDownloader {
	return &DefaultImageDownloader{}
}

func (d *DefaultImageDownloader) DownloadByUrl(ctx context.Context, url string) (*DownloadedImage, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, ErrRequest.Error())
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	downloadedImage := &DownloadedImage{img: img, headers: resp.Header}

	return downloadedImage, nil
}
