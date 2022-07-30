package queue

import (
	"context"
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		q := NewQueue(3)
		q.Enqueue(context.Background(), 1)
		q.Enqueue(context.Background(), 2)
		ok := q.TryEnqueue(3)
		assert(ok, t)

		ok = q.TryEnqueue(4)
		assert(!ok, t)
		v, ok := q.Dequeue(context.Background())
		assert(ok, t)
		assert(v == 1, t)
		v, ok = q.Dequeue(context.Background())
		assert(ok, t)
		assert(v == 2, t)
		v, ok = q.TryDequeue()
		assert(ok, t)
		assert(v == 3, t)
		v, ok = q.TryDequeue()
		assert(!ok, t)
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		go func() {
			<-time.After(time.Second * 3)
			cancel()
		}()
		v, ok = q.Dequeue(ctx)
		assert(!ok, t)
	})
}

func assert(b bool, t *testing.T) {
	t.Helper()
	if !b {
		t.Errorf("")
	}
}
