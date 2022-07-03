package lock

import "sync"

type ConcurrentMap struct {
	sync.RWMutex
	m map[interface{}]interface{}
}

func New() *ConcurrentMap {
	return &ConcurrentMap{
		m: make(map[interface{}]interface{}),
	}
}

func (cm *ConcurrentMap) Get(k interface{}) (interface{}, bool) {
	cm.RLock()
	defer cm.RUnlock()

	v, exist := cm.m[k]
	return v, exist
}

func (cm *ConcurrentMap) Set(k interface{}, v interface{}) {
	cm.Lock()
	defer cm.Unlock()

	cm.m[k] = v
}

func (cm *ConcurrentMap) Delete(k interface{}) {
	cm.Lock()
	defer cm.Unlock()

	delete(cm.m, k)
}

func (cm *ConcurrentMap) Len() int {
	cm.RLock()
	defer cm.RUnlock()

	return len(cm.m)
}

func (cm *ConcurrentMap) Range(f func(k, v interface{}) bool) {
	cm.RLock()
	defer cm.RUnlock()

	for k, v := range cm.m {
		if !f(k, v) {
			return
		}
	}
}
