package sequential

import (
	"sync"
	"testing"
)

func TestH2o(t *testing.T) {
	h := newH2o()
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 6; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				h.h()
			}
			wg.Done()
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 600; i++ {
				h.o()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
