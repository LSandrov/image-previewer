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

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(2)
		c.Set("a", 1)
		c.Set("b", 2)

		hasValue1, ok := c.Get("a")
		require.Equal(t, 1, hasValue1)
		require.True(t, ok)

		hasValue2, ok := c.Get("b")
		require.Equal(t, 2, hasValue2)
		require.True(t, ok)

		c.Clear()

		clearedValue1, ok := c.Get("a")
		require.Nil(t, clearedValue1)
		require.False(t, ok)

		clearedValue2, ok := c.Get("b")
		require.Nil(t, clearedValue2)
		require.False(t, ok)
	})

	t.Run("push logic", func(t *testing.T) {
		c := NewCache(2)
		c.Set("a", 1)
		c.Set("b", 2)

		hasValue1, ok := c.Get("a")
		require.Equal(t, 1, hasValue1)
		require.True(t, ok)

		hasValue2, ok := c.Get("b")
		require.Equal(t, 2, hasValue2)
		require.True(t, ok)

		c.Set("c", 3)
		hasValue3, ok := c.Get("c")
		require.Equal(t, 3, hasValue3)
		require.True(t, ok)

		clearedValue1, ok := c.Get("a")
		require.Nil(t, clearedValue1)
		require.False(t, ok)
	})

	t.Run("push last logic", func(t *testing.T) {
		c := NewCache(3)
		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)

		hasValue2, ok := c.Get("b")
		require.Equal(t, 2, hasValue2)
		require.True(t, ok)

		hasValue1, ok := c.Get("a")
		require.Equal(t, 1, hasValue1)
		require.True(t, ok)

		hasValue3, ok := c.Get("c")
		require.Equal(t, 3, hasValue3)
		require.True(t, ok)

		_, ok = c.Get("a")
		require.True(t, ok)
		_, ok = c.Get("c")
		require.True(t, ok)

		c.Set("d", 4)

		clearedValue1, ok := c.Get("b")
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
			c.Set(strconv.Itoa(i), i)
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
