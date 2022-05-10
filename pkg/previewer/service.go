package previewer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image-previewer/pkg/cache"
	"image/jpeg"
	"io"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/rs/zerolog"
)

// Service сервис для обрезки изображений.
type Service interface {
	// Fill скачивает изображение, обрезает его. Данные кешируются.
	Fill(params *FillParams) (*FillResponse, error)
}

type DefaultService struct {
	l               zerolog.Logger
	downloader      ImageDownloader
	resizedCache    cache.Cache
	downloadedCache cache.Cache
}

func NewDefaultService(
	l zerolog.Logger,
	downloader ImageDownloader,
	resizedCache cache.Cache,
	downloadedCache cache.Cache,
) Service {
	return &DefaultService{
		l:               l,
		downloader:      downloader,
		resizedCache:    resizedCache,
		downloadedCache: downloadedCache,
	}
}

func (svc *DefaultService) Fill(params *FillParams) (*FillResponse, error) {
	ctx, cancel := context.WithTimeout(params.ctx, time.Duration(5)*time.Second)
	defer cancel()

	cacheKey := svc.resizedCache.MakeCacheKeyResizes(params.width, params.height, params.url)
	cached, ok := svc.resizedCache.Get(cacheKey)

	if ok {
		fillResponse := NewFillResponse(cached.Img, cached.Header)
		return fillResponse, nil
	}

	cacheKeyDownloaded := svc.downloadedCache.MakeCacheKeyDownloaded(params.url)
	cachedImage, ok := svc.downloadedCache.Get(cacheKeyDownloaded)

	var downloaded *DownloadedImage

	if ok {
		downloaded = &DownloadedImage{img: cachedImage.Img, headers: cachedImage.Header}
	} else {
		downloaded, err := svc.downloader.DownloadByURL(ctx, params.url, params.headers)
		if err != nil {
			svc.l.Err(err).Msg("Невозможно загрузить изображение")
			return nil, err
		}

		go svc.downloadedCache.Set(&cache.Item{
			Key:    cacheKeyDownloaded,
			Img:    downloaded.img,
			Header: downloaded.headers,
		})
	}

	if downloaded == nil {
		return nil, errors.New("ошибка при загрузке изображения")
	}

	resizedImg, err := svc.resize(downloaded.img, params.width, params.height)
	if err != nil {
		svc.l.Err(err).Msg("Невозможно обрезать изображение")
		return nil, err
	}

	go svc.resizedCache.Set(&cache.Item{
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
