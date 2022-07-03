package lock

import (
	"fmt"
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	for i := 0; i < 20; i++ {
		wg := sync.WaitGroup{}
		parallel := 10000
		wg.Add(parallel)
		l := NewLock()
		v := 0

		for i := 0; i < parallel; i++ {
			go func() {
				for i := 0; i < 10; i++ {
					if !l.TryLock() {
						l.Lock()
					}
					v++
					l.Unlock()
				}
				wg.Done()
			}()
		}

		wg.Wait()
		fmt.Println(v)
		if v != parallel*10 {
			t.Error()
		}
	}
}
