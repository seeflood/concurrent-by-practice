package concurrent

import (
	"fmt"
	"sync"
	"time"
)

func Crawl(seed []string) <-chan string {
	resultIn, resultOut := makeChan()
	qIn, qOut := makeChan()
	var visited sync.Map
	var wg sync.WaitGroup
	wg.Add(k)

	for i := 0; i < k; i++ {
		go func(idx int) {
			for {
				var url string
				// check closed
				select {
				case <-time.After(500 * time.Millisecond):
					fmt.Printf("goroutine %v exited\n", idx)
					wg.Done()
					return
				case url = <-qOut:
				}
				// crawl
				r := doCrawl(url)
				resultIn <- r
				// get related urls
				urls := getUrls(url)
				for _, u := range urls {
					if !needCrawl(u) {
						continue
					}
					if _, ok := visited.Load(u); ok {
						continue
					}
					visited.Store(u, struct{}{})
					qIn <- u
				}
			}
		}(i)
	}
	for i := 0; i < len(seed); i++ {
		visited.Store(seed[i], struct{}{})
		qIn <- seed[i]
	}
	go func() {
		wg.Wait()
		fmt.Println("try to close resultIn")
		close(resultIn)
	}()
	return resultOut
}

func doCrawl(url string) string {
	time.Sleep(200 * time.Millisecond)
	return url + " result"
}

func getUrls(url string) []string {
	time.Sleep(200 * time.Millisecond)
	return []string{url, "b", "c"}
}

func needCrawl(url string) bool {
	return true
}

const chanSize = 1

func makeChan() (chan<- string, <-chan string) {
	in, out := make(chan string, chanSize), make(chan string)
	cache := make([]string, 0)
	go func() {
		for {
			v, ok := <-in
			if !ok {
				// closed
				fmt.Println("try to close out...")
				close(out)
				return
			}
			cache = append(cache, v)
			// move cache to out bound channel until it's empty
			for len(cache) > 0 {
				select {
				case out <- cache[0]:
					cache = cache[1:]
				case v := <-in:
					cache = append(cache, v)
				}
			}
		}
	}()
	return in, out
}
