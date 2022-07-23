package my_sync

import "sync/atomic"

type Once2 struct {
	started uint32
	// commit flag
	commited uint32
}

func (o *Once2) Do(f func()) {
	// check result
	if atomic.LoadUint32(&o.commited) == 1 {
		return
	}
	// try lock
	if atomic.CompareAndSwapUint32(&o.started, 0, 1) {
		// got lock
		defer atomic.StoreUint32(&o.commited, 1)
		f()
		return
	}
	// wait until committed
	for atomic.LoadUint32(&o.commited) != 1 {
	}
}
