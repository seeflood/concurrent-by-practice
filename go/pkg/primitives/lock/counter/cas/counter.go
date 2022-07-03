package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const k = 100000

func main() {
	var cnt uint64
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < k; i++ {
				atomic.AddUint64(&cnt, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(cnt)
}
