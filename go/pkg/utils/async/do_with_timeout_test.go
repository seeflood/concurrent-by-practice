package async

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDoWithTimeout(t *testing.T) {
	t.Run("not printed", func(t *testing.T) {
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
		DoWithTimeout(ctx, func() {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("A printed!")
		})
	})
}
