package previewer

import (
	"context"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ErrTimeout     = errors.New("таймаут загрузки изображения")
	ErrRequest     = errors.New("ошибка при отправке запроса на загрузку")
	ErrImgValidate = errors.New("файл не является изображением")
)

type ImageDownloader interface {
	DownloadByUrl(ctx context.Context, url string, headers map[string][]string) (*DownloadedImage, error)
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

func (d *DefaultImageDownloader) DownloadByUrl(ctx context.Context, url string, headers map[string][]string) (*DownloadedImage, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, errors.Wrap(err, ErrRequest.Error())
	}

	req.Header = headers

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := d.validate(img); err != nil {
		return nil, err
	}

	downloadedImage := &DownloadedImage{img: img, headers: resp.Header}

	if err := resp.Body.Close(); err != nil {
		return nil, err
	}

	return downloadedImage, nil
}

func (d *DefaultImageDownloader) validate(img []byte) error {
	if len(img) == 0 {
		return ErrRequest
	}

	allowedFormats := map[string]string{
		"\xff\xd8\xff":      "image/jpeg",
		"\x89PNG\r\n\x1a\n": "image/png",
		"GIF87a":            "image/gif",
		"GIF89a":            "image/gif",
	}

	imgStr := string(img)
	for format, _ := range allowedFormats {
		if strings.HasPrefix(imgStr, format) {
			return nil
		}
	}

	return ErrImgValidate
}
