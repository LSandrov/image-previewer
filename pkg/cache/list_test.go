package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		require.Equal(t, 40, l.Front().Next.Next.Value)

		l.MoveToFront(l.Back()) // [70, 80, 60, 40, 10, 30, 50]
		require.Nil(t, l.Back().Next)

		l.MoveToFront(l.Back().Prev) // [30, 70, 80, 60, 40, 10, 50]
		require.Equal(t, 80, l.Front().Next.Next.Value)
		require.Equal(t, 10, l.Back().Prev.Value)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{30, 70, 80, 60, 40, 10, 50}, elems)
	})

	t.Run("custom", func(t *testing.T) {
		l := NewList()

		for i, v := range [...]int{1, 2, 3, 4, 5, 6, 7, 8} {
			if i < 5 {
				l.PushBack(v)
			} else {
				l.PushFront(v)
			}
		}

		require.Equal(t, 8, l.Len())
		require.Equal(t, 8, l.Front().Value)
		require.Equal(t, 6, l.Back().Value)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{8, 7, 6}, elems)
	})

	t.Run("remove_without_next", func(t *testing.T) {
		l := NewList()

		for i, v := range [...]int{1, 2, 3, 4, 5, 6, 7, 8} {
			if i < 5 {
				l.PushBack(v)
			} else {
				l.PushFront(v)
			}
		}

		require.Equal(t, 8, l.Len())
		require.Equal(t, 8, l.Front().Value)
		require.Equal(t, 6, l.Back().Value)

		backPrev := l.Back().Prev
		l.Remove(l.Back())

		require.Equal(t, l.Back(), backPrev)
	})

	t.Run("remove_without_back", func(t *testing.T) {
		l := NewList()

		for i, v := range [...]int{1, 2, 3, 4, 5, 6, 7, 8} {
			if i < 5 {
				l.PushBack(v)
			} else {
				l.PushFront(v)
			}
		}

		require.Equal(t, 8, l.Len())
		require.Equal(t, 8, l.Front().Value)
		require.Equal(t, 6, l.Back().Value)

		frontNext := l.Front().Next
		l.Remove(l.Front())

		require.Equal(t, l.Front(), frontNext)
	})
}
