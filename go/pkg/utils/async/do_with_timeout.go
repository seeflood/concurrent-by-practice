package async

import (
	"context"
)

func DoWithTimeout(ctx context.Context, f func()) {
	out := make(chan struct{})
	go func() {
		f()
		out <- struct{}{}
	}()
	select {
	case <-out:
	case <-ctx.Done():
	}
}
