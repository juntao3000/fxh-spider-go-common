package baseComponent

import (
	"sync"
	"time"
)

type CachedConcurrentSet struct {
	itemMap       map[string]time.Time
	cacheDuration time.Duration
	rwMutex       sync.RWMutex
}

func NewCachedConcurrentSet(cacheDuration time.Duration) *CachedConcurrentSet {
	return &CachedConcurrentSet{
		itemMap:       make(map[string]time.Time),
		cacheDuration: cacheDuration,
	}
}

func (m *CachedConcurrentSet) checkExpired(cacheTime time.Time) bool {
	return m.cacheDuration > 0 && time.Since(cacheTime) > m.cacheDuration
}

func (m *CachedConcurrentSet) CleanExpired() {
	m.rwMutex.Lock()
	keyList := make([]string, 0)
	for key, item := range m.itemMap {
		if m.checkExpired(item) {
			keyList = append(keyList, key)
			continue
		}
	}
	for _, key := range keyList {
		delete(m.itemMap, key)
	}
	m.rwMutex.Unlock()
}

func (m *CachedConcurrentSet) Range(f func(key string) bool) {
	m.rwMutex.RLock()
	for key, item := range m.itemMap {
		if m.checkExpired(item) {
			continue
		}

		if !f(key) {
			break
		}
	}
	m.rwMutex.RUnlock()
}

func (m *CachedConcurrentSet) Count() int {
	count := 0
	m.rwMutex.RLock()
	for _, item := range m.itemMap {
		if m.checkExpired(item) {
			continue
		}
		count++
	}
	m.rwMutex.RUnlock()
	return count
}

func (m *CachedConcurrentSet) Delete(key string) {
	m.rwMutex.Lock()
	delete(m.itemMap, key)
	m.rwMutex.Unlock()
}

func (m *CachedConcurrentSet) Exists(key string) bool {
	m.rwMutex.RLock()
	item, ok := m.itemMap[key]
	if !ok {
		m.rwMutex.RUnlock()
		return false
	}

	expired := m.checkExpired(item)
	if expired {
		m.rwMutex.RUnlock()
		return false
	}

	m.rwMutex.RUnlock()
	return true
}

func (m *CachedConcurrentSet) Add(key string) {
	m.rwMutex.Lock()
	m.itemMap[key] = time.Now()
	m.rwMutex.Unlock()
}
