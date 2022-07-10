package semaphore

import "github.com/seeflood/concurrent_challenges/go/pkg/utils/queue"

type Semaphore struct {
	q *queue.UnboundedQueue
}

func NewSemaphore(initialPermits uint32) *Semaphore {
	s := &Semaphore{
		q: queue.NewUnboundedQueue(),
	}
	var i uint32 = 0
	for ; i < initialPermits; i++ {
		s.Release()
	}
	return s
}

func (s *Semaphore) Acquire() {
	s.q.Deque()
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case <-s.q.DequeChannel():
		return true
	default:
		return false
	}
}

func (s *Semaphore) Release() {
	s.q.Enque(struct{}{})
}
