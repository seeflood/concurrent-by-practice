package my_sync

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	// commit flag
	flag  uint32
	mutex sync.Mutex
}

func (o *Once) Do(f func()) {
	// DCL
	if atomic.LoadUint32(&o.flag) == 1 {
		return
	}
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.flag == 1 {
		return
	}
	defer atomic.StoreUint32(&o.flag, 1)
	f()
}
