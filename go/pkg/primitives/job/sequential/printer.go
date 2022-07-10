package sequential

import (
	"fmt"
	"sync"
)

const k = 3

// 随机调度, 缺点是没有按题目要求"各个线程依次执行"
func printInOrderRandom() {
	v := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(k)
	// start workers
	for i := 0; i < k; i++ {
		go func(id int) {
			for {
				mutex.Lock()
				if v > 100 {
					mutex.Unlock()
					wg.Done()
					return
				}
				fmt.Printf("goroutine %v: %v\n", id, v)
				v++
				mutex.Unlock()
			}
		}(i)
	}
	// wait
	wg.Wait()
	fmt.Println("All printed!")
}

// 各个线程依次执行
func printInOrder() {
	closed := make(chan struct{})
	// init outs channels
	outs := make([]chan int, k)
	for i := 0; i < k; i++ {
		outs[i] = make(chan int, 1)
	}
	// start workers
	for i := 0; i < k; i++ {
		go func(id int) {
			for {
				prev := (id + k - 1) % k
				// wait for prev
				v, ok := <-outs[prev]
				if !ok {
					return
				}
				// print
				fmt.Printf("goroutine %v: %v\n", id, v)
				if v == 100 {
					close(closed)
				} else {
					// notify next goroutine
					outs[id] <- v + 1
				}
			}
		}(i)
	}
	outs[k-1] <- 0
	<-closed
	for i := 0; i < k; i++ {
		close(outs[i])
	}
	fmt.Println("All printed!")
}
