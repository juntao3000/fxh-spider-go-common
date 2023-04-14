package baseComponent

import (
	"sync"
)

type ConcurrentSet struct {
	itemMap map[string]struct{}
	rwMutex sync.RWMutex
}

func NewConcurrentSet() *ConcurrentSet {
	return &ConcurrentSet{
		itemMap: make(map[string]struct{}),
	}
}

func (m *ConcurrentSet) Range(f func(key string) bool) {
	m.rwMutex.RLock()
	for key := range m.itemMap {
		if !f(key) {
			break
		}
	}
	m.rwMutex.RUnlock()
}

func (m *ConcurrentSet) Count() int {
	count := 0
	m.rwMutex.RLock()
	count = len(m.itemMap)
	m.rwMutex.RUnlock()
	return count
}

func (m *ConcurrentSet) Delete(key string) {
	m.rwMutex.Lock()
	delete(m.itemMap, key)
	m.rwMutex.Unlock()
}

func (m *ConcurrentSet) Exists(key string) bool {
	ok := false
	m.rwMutex.RLock()
	_, ok = m.itemMap[key]
	m.rwMutex.RUnlock()
	return ok
}

func (m *ConcurrentSet) Add(key string) {
	m.rwMutex.Lock()
	m.itemMap[key] = struct{}{}
	m.rwMutex.Unlock()
}
