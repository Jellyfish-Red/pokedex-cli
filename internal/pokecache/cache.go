package pokecache

import (
	"sync"
	"time"
)

type APICache struct {
	entries map[string]CacheEntry
	cacheLock sync.Mutex

	reapInterval time.Duration
}

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) APICache {
	cache := APICache{
		entries: map[string]CacheEntry{},
		cacheLock: sync.Mutex{},
		reapInterval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (cache *APICache) Add(key string, val []byte) {
	cache.cacheLock.Lock()
	cache.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	cache.cacheLock.Unlock()
}

func (cache *APICache) Get(key string) ([]byte, bool) {
	cache.cacheLock.Lock()
	elem, ok := cache.entries[key]
	cache.cacheLock.Unlock()

	if ok {
		return elem.val, ok
	}

	return nil, false
}

func (cache *APICache) reapLoop() {
	ticker := time.NewTicker(cache.reapInterval)
	for {
		<-ticker.C
		for name, entry := range cache.entries {
			if time.Since(entry.createdAt) >= cache.reapInterval {
				delete(cache.entries, name)
			}
		}
	}
}