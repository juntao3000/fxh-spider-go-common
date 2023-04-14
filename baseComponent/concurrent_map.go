package baseComponent

import (
	"sync"
)

type ConcurrentMap struct {
	itemMap map[string]any
	rwMutex sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		itemMap: make(map[string]any),
	}
}

func (m *ConcurrentMap) Range(f func(key string, value any) bool) {
	m.rwMutex.RLock()
	for key, item := range m.itemMap {
		if !f(key, item) {
			break
		}
	}
	m.rwMutex.RUnlock()
}

func (m *ConcurrentMap) Count() int {
	count := 0
	m.rwMutex.RLock()
	count = len(m.itemMap)
	m.rwMutex.RUnlock()
	return count
}

func (m *ConcurrentMap) Get(key string) (any, bool) {
	m.rwMutex.RLock()
	item, ok := m.itemMap[key]
	m.rwMutex.RUnlock()
	return item, ok
}

func (m *ConcurrentMap) Delete(key string) {
	m.rwMutex.Lock()
	delete(m.itemMap, key)
	m.rwMutex.Unlock()
}

func (m *ConcurrentMap) Exists(key string) bool {
	m.rwMutex.RLock()
	_, ok := m.itemMap[key]
	m.rwMutex.RUnlock()
	return ok
}

func (m *ConcurrentMap) Set(key string, value any) {
	m.rwMutex.Lock()
	m.itemMap[key] = value
	m.rwMutex.Unlock()
}
