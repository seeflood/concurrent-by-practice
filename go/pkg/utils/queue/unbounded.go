package queue

type UnboundedQueue struct {
	in  chan<- interface{}
	out <-chan interface{}
}

func NewUnboundedQueue() *UnboundedQueue {
	in, out := make(chan interface{}), make(chan interface{})
	// start consumer
	go func() {
		cache := make([]interface{}, 0)
		for {
			// move the data from `in` to cache
			e, ok := <-in
			if !ok {
				close(out)
				return
			}
			cache = append(cache, e)
			// consume all the cache while watching the in bound channel
			for len(cache) > 0 {
				select {
				case out <- cache[0]:
					cache = cache[1:]
				case e, ok := <-in:
					if ok {
						cache = append(cache, e)
						break
					}
					// in bound channel closed
					// drain the cache
					for _, e = range cache {
						out <- e
					}
					close(out)
					return
				}
			}
		}
	}()
	return &UnboundedQueue{
		in:  in,
		out: out,
	}
}

func (q *UnboundedQueue) Close() {
	close(q.in)
}

func (q *UnboundedQueue) Enque(ele interface{}) {
	q.in <- ele
}

func (q *UnboundedQueue) Deque() (interface{}, bool) {
	v, ok := <-q.out
	return v, ok
}

func (q *UnboundedQueue) DequeChannel() <-chan interface{} {
	return q.out
}
