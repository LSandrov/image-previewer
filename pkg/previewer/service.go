package previewer

import (
	"bytes"
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/rs/zerolog"
	"image-previewer/pkg/cache"
	"image/jpeg"
	"io"
	"os"
	"time"
)

// Service сервис для обрезки изображений
type Service interface {
	// Fill скачивает изображение, обрезает его. Данные кешируются
	Fill(ctx context.Context, width, height int, imgURL string) (*FillResponse, error)
}

type DefaultService struct {
	l          zerolog.Logger
	downloader ImageDownloader
	cache      cache.Cache
}

func NewDefaultService(l zerolog.Logger, downloader ImageDownloader, c cache.Cache) Service {
	return &DefaultService{l: l, downloader: downloader, cache: c}
}

func (svc *DefaultService) Fill(ctx context.Context, width, height int, imgURL string) (*FillResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(1000)*time.Second)
	defer cancel()

	cacheKey := svc.cache.MakeCacheKeyResizes(width, height, imgURL)
	cached, ok := svc.cache.Get(cacheKey)

	if ok {
		fillResponse := NewFillResponse(cached.Img, cached.Header)
		return fillResponse, nil
	}

	cacheKeyDownloaded := svc.cache.MakeCacheKeyDownloaded(imgURL)
	cachedImage, ok := svc.cache.Get(cacheKeyDownloaded)

	if ok {
		fillResponse := NewFillResponse(cachedImage.Img, cachedImage.Header)

		return fillResponse, nil
	}

	downloaded, err := svc.downloader.DownloadByUrl(ctx, imgURL)
	if err != nil {
		svc.l.Err(err).Msg("Невозможно загрузить изображение")
		return nil, err
	}

	svc.cache.Set(cache.Item{
		Key:    cacheKeyDownloaded,
		Img:    downloaded.img,
		Header: downloaded.headers,
	})

	resizedImg, err := svc.resize(downloaded.img, width, height)
	if err != nil {
		svc.l.Err(err).Msg("Невозможно обрезать изображение")
		return nil, err
	}

	svc.cache.Set(cache.Item{
		Key:    cacheKey,
		Img:    resizedImg,
		Header: downloaded.headers,
	})

	fillResponse := NewFillResponse(resizedImg, downloaded.headers)

	return fillResponse, nil
}

func (svc *DefaultService) resize(img []byte, width, height int) ([]byte, error) {
	tmpImgName := fmt.Sprintf("%d.jpg", time.Now().Unix())

	file, err := os.Create(tmpImgName)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			svc.l.Err(err).Msg("ошибка при работе с файлом")
		}
	}(file)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			svc.l.Err(err).Msg("ошибка при удалении временного файла")
		}
	}(tmpImgName)

	_, err = io.Copy(file, bytes.NewReader(img))
	if err != nil {
		svc.l.Err(err).Msg(err.Error())
	}

	src, err := imaging.Open(tmpImgName)
	if err != nil {
		return nil, err
	}

	resized := imaging.Resize(src, width, height, imaging.Lanczos)
	imgBuffer := new(bytes.Buffer)
	err = jpeg.Encode(imgBuffer, resized, nil)

	return imgBuffer.Bytes(), err
}
