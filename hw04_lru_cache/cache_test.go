package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCacheSmallSize(t *testing.T) {
	t.Run("size is one", func(t *testing.T) {
		c := NewCache(1)

		ok := c.Set("aaa", 10)
		require.False(t, ok)

		_, ok = c.Get("aaa")
		require.True(t, ok)

		ok = c.Set("aaa", 20)
		require.True(t, ok)

		_, ok = c.Get("aaa")
		require.True(t, ok)

		ok = c.Set("bbb", 30)
		require.False(t, ok)

		_, ok = c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.True(t, ok)
	})
}

func TestCachePushout(t *testing.T) {
	// the previous one was "push out"
	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		ok := c.Set("aaa", 10) // |{aaa,10}|  |  |
		require.False(t, ok)

		ok = c.Set("bbb", 20) // |{bbb,20}|{aaa,10}|  |
		require.False(t, ok)

		ok = c.Set("ccc", 30) // |{ccc,30}|{bbb,20}|{aaa,10}|
		require.False(t, ok)

		ok = c.Set("aaa", 30) // |{aaa,30}|{ccc,30}|{bbb,20}|
		require.True(t, ok)

		ok = c.Set("bbb", 20) // |{bbb,20}|{aaa,30}|{ccc,30}|
		require.True(t, ok)

		ok = c.Set("ccc", 10) // |{ccc,10}|{bbb,20}|{aaa,30}|
		require.True(t, ok)

		ok = c.Set("ddd", 40) // |{ddd,40}|{ccc,10}|{bbb,20}|	-> {aaa,20}
		require.False(t, ok)

		_, ok = c.Get("aaa") // |{ddd,40}|{ccc,10}|{bbb,20}|	 x {aaa,20}
		require.False(t, ok)

		ok = c.Set("eee", 50) // |{eee,50}|{ddd,40}|{ccc,10}|	-> {bbb,20}
		require.False(t, ok)

		_, ok = c.Get("bbb") // |{eee,50}|{ddd,40}|{ccc,10}|	 x {bbb,20}
		require.False(t, ok)

		ok = c.Set("fff", 50) // |{fff,50}|{eee,50}|{ddd,40}|	-> {ccc,10}
		require.False(t, ok)

		_, ok = c.Get("ccc") // |{fff,50}|{eee,50}|{ddd,40}|	 x {ccc,10}
		require.False(t, ok)
	})

	t.Run("push the oldest item out", func(t *testing.T) {
		c := NewCache(5)

		ok := c.Set("aaa", 10)
		require.False(t, ok)
		ok = c.Set("bbb", 20)
		require.False(t, ok)
		ok = c.Set("ccc", 30)
		require.False(t, ok)
		ok = c.Set("ddd", 40)
		require.False(t, ok)
		ok = c.Set("eee", 50)
		require.False(t, ok)

		ok = c.Set("aaa", 11)
		require.True(t, ok)
		ok = c.Set("bbb", 21)
		require.True(t, ok)
		ok = c.Set("ccc", 31)
		require.True(t, ok)
		ok = c.Set("eee", 51)
		require.True(t, ok)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 11, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 21, val)

		val, ok = c.Get("ccc")
		require.True(t, ok)
		require.Equal(t, 31, val)

		ok = c.Set("fff", 60)
		require.False(t, ok)

		val, ok = c.Get("ddd")
		require.False(t, ok)
		require.Equal(t, nil, val)
	})
}

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

	t.Run("clear", func(t *testing.T) {
		c := NewCache(10)

		for i := 0; i < 20; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}

		c.Clear()

		for i := 0; i < 20; i++ {
			_, ok := c.Get(Key(strconv.Itoa(i)))
			require.False(t, ok)
		}

		for i := 0; i < 20; i++ {
			ok := c.Set(Key(strconv.Itoa(i)), i)
			require.False(t, ok)
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	// t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
