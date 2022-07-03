package main

import (
	"fmt"
	"sync"
)

const k = 100000

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < k; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

type Counter struct {
	sync.Mutex
	cnt uint64
}

func (c *Counter) Incr() {
	c.Lock()
	defer c.Unlock()

	c.cnt++
}

func (c *Counter) Count() uint64 {
	c.Lock()
	defer c.Unlock()

	return c.cnt
}
