package cache

type Cache interface {
	Get(key string) (val []byte, err error)
	Set(key string, val []byte) (err error)
}
