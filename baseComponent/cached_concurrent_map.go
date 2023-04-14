package baseComponent

import (
	"sync"
	"time"
)

type cachedConcurrentMapItem struct {
	value     any
	cacheTime time.Time
}

type CachedConcurrentMap struct {
	itemMap       map[string]*cachedConcurrentMapItem
	cacheDuration time.Duration
	rwMutex       sync.RWMutex
}

func NewCachedConcurrentMap(cacheDuration time.Duration) *CachedConcurrentMap {
	return &CachedConcurrentMap{
		itemMap:       make(map[string]*cachedConcurrentMapItem),
		cacheDuration: cacheDuration,
	}
}

func (m *CachedConcurrentMap) checkExpired(item *cachedConcurrentMapItem) bool {
	return m.cacheDuration > 0 && time.Since(item.cacheTime) > m.cacheDuration
}

func (m *CachedConcurrentMap) CleanExpired() {
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

func (m *CachedConcurrentMap) Range(f func(key string, value any) bool) {
	m.rwMutex.RLock()
	for key, item := range m.itemMap {
		if m.checkExpired(item) {
			continue
		}

		if !f(key, item.value) {
			break
		}
	}
	m.rwMutex.RUnlock()
}

func (m *CachedConcurrentMap) Count() int {
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

func (m *CachedConcurrentMap) Get(key string) (any, bool) {
	m.rwMutex.RLock()
	item, ok := m.itemMap[key]
	if !ok {
		m.rwMutex.RUnlock()
		return nil, false
	}

	expired := m.checkExpired(item)
	if expired {
		m.rwMutex.RUnlock()
		return nil, false
	}

	m.rwMutex.RUnlock()
	return item.value, true
}

func (m *CachedConcurrentMap) Delete(key string) {
	m.rwMutex.Lock()
	delete(m.itemMap, key)
	m.rwMutex.Unlock()
}

func (m *CachedConcurrentMap) Exists(key string) bool {
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

func (m *CachedConcurrentMap) Set(key string, value any) {
	m.rwMutex.Lock()
	m.itemMap[key] = &cachedConcurrentMapItem{
		cacheTime: time.Now(),
		value:     value,
	}
	m.rwMutex.Unlock()
}
