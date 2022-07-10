package semaphore

type BoundedSemaphore struct {
	initialPermits int
	ch             chan struct{}
}

func NewBoundedSemaphore(initialPermits int) *BoundedSemaphore {
	bs := &BoundedSemaphore{
		initialPermits: initialPermits,
		ch:             make(chan struct{}, initialPermits),
	}
	for i := 0; i < initialPermits; i++ {
		bs.Release()
	}
	return bs
}

func (bs *BoundedSemaphore) Acquire() {
	<-bs.ch
}

func (bs *BoundedSemaphore) TryAcquire() bool {
	select {
	case <-bs.ch:
		return true
	default:
		return false
	}
}

func (bs *BoundedSemaphore) Release() {
	bs.ch <- struct{}{}
}
