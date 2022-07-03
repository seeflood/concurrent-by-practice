package semaphore

import "github.com/seeflood/concurrent_challenges/go/pkg/utils/queue"

type Semaphore struct {
	q *queue.UnboundedQueue
}

func NewSemaphore(cap uint32) *Semaphore {
	return &Semaphore{
		q: queue.NewUnboundedQueue(),
	}
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
