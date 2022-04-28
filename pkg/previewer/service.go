package previewer

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"image-previewer/pkg/cache"
)

type Service interface {
	Fill(ctx context.Context, width, height int, imgUrl string) (img []byte, err error)
}

type DefaultService struct {
	l          zerolog.Logger
	downloader ImageDownloader
	cache      cache.Cache
}

func NewDefaultService(l zerolog.Logger, downloader ImageDownloader, c cache.Cache) Service {
	return &DefaultService{l: l, downloader: downloader, cache: c}
}

func (svc *DefaultService) Fill(ctx context.Context, width, height int, imgUrl string) (img []byte, err error) {
	cacheKey := svc.makeCacheKey(width, height, imgUrl)
	cachedImg, ok := svc.cache.Get(cacheKey)

	if ok {
		img = cachedImg.([]byte)
	}

	img, err = svc.downloader.DownloadByUrl(ctx, imgUrl)
	if err != nil {
		svc.l.Error().Msg("Невозможно загрузить изображение")
		return nil, err
	}

	return img, nil
}

func (svc *DefaultService) makeCacheKey(width, height int, url string) string {
	return fmt.Sprintf("%d_%d_%s", width, height, url)
}
