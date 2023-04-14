package baseComponent

import (
	"time"
)

type cachedMapItem struct {
	value     any
	cacheTime time.Time
}

type CachedMap struct {
	itemMap       map[string]*cachedMapItem
	cacheDuration time.Duration
}

func NewCachedMap(cacheDuration time.Duration) *CachedMap {
	return &CachedMap{
		itemMap:       make(map[string]*cachedMapItem),
		cacheDuration: cacheDuration,
	}
}

func (m *CachedMap) checkExpired(item *cachedMapItem) bool {
	return m.cacheDuration > 0 && time.Since(item.cacheTime) > m.cacheDuration
}

func (m *CachedMap) CleanExpired() {
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
}

func (m *CachedMap) Range(f func(key string, value any) bool) {
	for key, item := range m.itemMap {
		if m.checkExpired(item) {
			continue
		}

		if !f(key, item.value) {
			break
		}
	}
}

func (m *CachedMap) Count() int {
	count := 0
	for _, item := range m.itemMap {
		if m.checkExpired(item) {
			continue
		}
		count++
	}
	return count
}

func (m *CachedMap) Get(key string) (any, bool) {
	item, ok := m.itemMap[key]
	if !ok {
		return nil, false
	}

	expired := m.checkExpired(item)
	if expired {
		return nil, false
	}

	return item.value, true
}

func (m *CachedMap) Delete(key string) {
	delete(m.itemMap, key)
}

func (m *CachedMap) Exists(key string) bool {
	item, ok := m.itemMap[key]
	if !ok {
		return false
	}
	return m.checkExpired(item)
}

func (m *CachedMap) Set(key string, value any) {
	m.itemMap[key] = &cachedMapItem{
		cacheTime: time.Now(),
		value:     value,
	}
}
