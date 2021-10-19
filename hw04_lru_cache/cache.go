package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		mu:       sync.RWMutex{},
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set - insert value to cache by key.
func (lr *lruCache) Set(key Key, value interface{}) bool {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	ci := cacheItem{key: string(key), value: value}
	if addr, ok := lr.items[key]; ok {
		addr.Value = ci
		lr.queue.MoveToFront(addr)
		return true
	}

	if lr.capacity == lr.queue.Len() {
		addr := lr.queue.Back()
		lr.queue.Remove(addr)
		delete(lr.items, Key(addr.Value.(cacheItem).key))
	}

	l := lr.queue.PushFront(ci)
	lr.items[key] = l

	return false
}

// Get - return value from cache by key.
func (lr *lruCache) Get(key Key) (interface{}, bool) {
	lr.mu.RLock()
	defer lr.mu.RUnlock()
	if addr, ok := lr.items[key]; ok {
		lr.queue.MoveToFront(addr)
		return addr.Value.(cacheItem).value, true
	}
	return nil, false
}

// Clear - remove all items from cache.
func (lr *lruCache) Clear() {
	lr.mu.Lock()
	defer lr.mu.Unlock()
	for i := range lr.items {
		delete(lr.items, i)
	}

	for i := lr.queue.Back(); i != nil; i = i.Prev {
		lr.queue.Remove(i)
	}
}
