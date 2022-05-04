package cache

type Cache interface {
	Get(key string) (*Item, bool)
	Set(i Item) bool
	MakeCacheKeyResizes(width, height int, url string) string
	MakeCacheKeyDownloaded(url string) string
	Clear()
}

type Item struct {
	Key    string
	Img    []byte
	Header map[string][]string
}
