package cache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		itemA := &Item{
			Key:    "aaa",
			Img:    []byte("aaa"),
			Header: make(map[string][]string),
		}

		itemB := &Item{
			Key:    "bbb",
			Img:    []byte("bbb"),
			Header: make(map[string][]string),
		}

		itemC := &Item{
			Key:    "ccc",
			Img:    []byte("ccc"),
			Header: make(map[string][]string),
		}

		wasInCache := c.Set(itemA)
		require.False(t, wasInCache)

		wasInCache = c.Set(itemB)
		require.False(t, wasInCache)

		val, ok := c.Get(itemA.Key)
		require.True(t, ok)
		require.Equal(t, itemA, val)

		val, ok = c.Get(itemB.Key)
		require.True(t, ok)
		require.Equal(t, itemB, val)

		wasInCache = c.Set(itemA)
		require.True(t, wasInCache)

		val, ok = c.Get(itemA.Key)
		require.True(t, ok)
		require.Equal(t, itemA, val)

		val, ok = c.Get(itemC.Key)
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(2)

		itemA := &Item{
			Key:    "aaa",
			Img:    []byte("aaa"),
			Header: make(map[string][]string),
		}

		itemB := &Item{
			Key:    "bbb",
			Img:    []byte("bbb"),
			Header: make(map[string][]string),
		}

		c.Set(itemA)
		c.Set(itemB)

		hasValue1, ok := c.Get(itemA.Key)
		require.Equal(t, itemA, hasValue1)
		require.True(t, ok)

		hasValue2, ok := c.Get(itemB.Key)
		require.Equal(t, itemB, hasValue2)
		require.True(t, ok)

		c.Clear()

		clearedValue1, ok := c.Get(itemA.Key)
		require.Nil(t, clearedValue1)
		require.False(t, ok)

		clearedValue2, ok := c.Get(itemB.Key)
		require.Nil(t, clearedValue2)
		require.False(t, ok)
	})

	t.Run("push logic", func(t *testing.T) {
		c := NewCache(2)

		itemA := &Item{
			Key:    "aaa",
			Img:    []byte("aaa"),
			Header: make(map[string][]string),
		}

		itemB := &Item{
			Key:    "bbb",
			Img:    []byte("bbb"),
			Header: make(map[string][]string),
		}

		itemC := &Item{
			Key:    "ccc",
			Img:    []byte("ccc"),
			Header: make(map[string][]string),
		}

		c.Set(itemA)
		c.Set(itemB)

		hasValue1, ok := c.Get(itemA.Key)
		require.Equal(t, itemA, hasValue1)
		require.True(t, ok)

		hasValue2, ok := c.Get(itemB.Key)
		require.Equal(t, itemB, hasValue2)
		require.True(t, ok)

		c.Set(itemC)
		hasValue3, ok := c.Get(itemC.Key)
		require.Equal(t, itemC, hasValue3)
		require.True(t, ok)

		clearedValue1, ok := c.Get(itemA.Key)
		require.Nil(t, clearedValue1)
		require.False(t, ok)
	})

	t.Run("push last logic", func(t *testing.T) {
		c := NewCache(3)

		itemA := &Item{
			Key:    "aaa",
			Img:    []byte("aaa"),
			Header: make(map[string][]string),
		}

		itemB := &Item{
			Key:    "bbb",
			Img:    []byte("bbb"),
			Header: make(map[string][]string),
		}

		itemC := &Item{
			Key:    "ccc",
			Img:    []byte("ccc"),
			Header: make(map[string][]string),
		}

		itemD := &Item{
			Key:    "ddd",
			Img:    []byte("ddd"),
			Header: make(map[string][]string),
		}

		c.Set(itemA)
		c.Set(itemB)
		c.Set(itemC)

		hasValue2, ok := c.Get(itemB.Key)
		require.Equal(t, itemB, hasValue2)
		require.True(t, ok)

		hasValue1, ok := c.Get(itemA.Key)
		require.Equal(t, itemA, hasValue1)
		require.True(t, ok)

		hasValue3, ok := c.Get(itemC.Key)
		require.Equal(t, itemC, hasValue3)
		require.True(t, ok)

		_, ok = c.Get(itemA.Key)
		require.True(t, ok)
		_, ok = c.Get(itemC.Key)
		require.True(t, ok)

		c.Set(itemD)

		clearedValue1, ok := c.Get(itemB.Key)
		require.Nil(t, clearedValue1)
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			item := &Item{
				Key:    strconv.Itoa(i),
				Img:    []byte("test"),
				Header: make(map[string][]string),
			}
			c.Set(item)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(strconv.Itoa(rand.Intn(1_000_000)))
		}
	}()

	wg.Wait()
}
