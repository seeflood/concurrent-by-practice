package concurrent

import (
	"fmt"
	"testing"
)

func TestCrawl(t *testing.T) {
	ch := Crawl([]string{
		"http://aaa",
		"http://b",
		"http://c",
		"http://d",
		"http://e",
		"http://f",
		"http://g",
		"http://h",
		"http://i",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
		"http://j",
	})
	for resp := range ch {
		fmt.Println(resp)
	}
	fmt.Println("all done!")
}
