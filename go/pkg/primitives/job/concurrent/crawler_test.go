package concurrent

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	urls := []string{
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
	}
	result := fetch(urls)
	fmt.Println(result)
	if len(urls) != len(result) {
		t.Error()
	}
}
