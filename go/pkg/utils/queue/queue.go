package queue

import "context"

type queue struct {
	ch chan interface{}
}

func NewQueue(cap int) Queue {
	return &queue{
		ch: make(chan interface{}, cap),
	}
}

type Queue interface {
	TryEnqueue(v interface{}) bool
	Enqueue(ctx context.Context, v interface{}) bool
	Dequeue(ctx context.Context) (interface{}, bool)
	TryDequeue() (interface{}, bool)
}

func (q *queue) TryEnqueue(v interface{}) bool {
	select {
	case q.ch <- v:
		return true
	default:
		return false
	}
}

func (q *queue) Enqueue(ctx context.Context, v interface{}) bool {
	select {
	case q.ch <- v:
		return true
	case <-ctx.Done():
		return false
	}
}

func (q *queue) Dequeue(ctx context.Context) (interface{}, bool) {
	select {
	case v := <-q.ch:
		return v, true
	case <-ctx.Done():
		return nil, false
	}
}

func (q *queue) TryDequeue() (interface{}, bool) {
	select {
	case v := <-q.ch:
		return v, true
	default:
		return nil, false
	}
}
