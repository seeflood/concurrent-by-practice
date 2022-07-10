package concurrent

import (
	"context"
	"fmt"
	"time"
)

func fetchWithTimeout(urls []string, timeout time.Duration) []string {
	n := len(urls)
	in, out := make(chan string, n), make(chan string, n)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// start consumer
	for i := 0; i < k; i++ {
		go func() {
			for {
				select {
				case url, ok := <-in:
					// all done
					if !ok {
						return
					}
					// call
					resp, ok := callWithCtx(ctx, url)
					if !ok {
						return
					}
					// put to out bound channel
					out <- resp
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// produce
	for i := 0; i < n; i++ {
		in <- urls[i]
	}
	close(in)
	// collect result
	result := make([]string, 0)
	for {
		select {
		case <-ctx.Done():
			return result
		case resp := <-out:
			result = append(result, resp)
		}
	}
	return result
}

func callWithCtx(ctx context.Context, url string) (string, bool) {
	select {
	case <-time.After(200 * time.Millisecond):
		return fmt.Sprintf("result of %v\n", url), true
	case <-ctx.Done():
		return "", false
	}
}
