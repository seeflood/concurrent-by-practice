package queue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestUnboundedQueue(t *testing.T) {
	q := NewUnboundedQueue()
	for i := 0; i < 100; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				q.Enque(i*10 + j)
			}
		}(i)
	}
	set := make([]bool, 100*10)
	var wg sync.WaitGroup
	wg.Add(100 * 10)
	//consume
	go func() {
		for {
			v, ok := q.Deque()
			if !ok {
				fmt.Println("The queue is closed!")
				break
			}
			fmt.Printf("consuming: %v \n", v)
			if set[v.(int)] {
				t.Error()
			}
			wg.Done()
			set[v.(int)] = true
			time.Sleep(100)
		}
	}()
	wg.Wait()
	q.Close()
	time.Sleep(1000)
	fmt.Println("Finished!")
}
