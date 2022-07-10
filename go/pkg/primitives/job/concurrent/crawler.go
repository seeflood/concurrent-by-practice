package concurrent

import (
	"time"
)

const k = 10

func fetch(urls []string) []string {
	n := len(urls)
	toVisit := make(chan string, n)
	result := make(chan string, n)
	// start consumer
	for i := 0; i < k; i++ {
		go func() {
			for url := range toVisit {
				r := call(url)
				result <- r
			}
		}()
	}
	// start producer
	for i := 0; i < n; i++ {
		toVisit <- urls[i]
	}
	// collect result
	r := make([]string, n)
	for i := 0; i < n; i++ {
		r[i] = <-result
	}
	return r
}

func call(url string) string {
	time.Sleep(200)
	return url
}
