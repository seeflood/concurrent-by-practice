package concurrent

import (
	"fmt"
	"testing"
	"time"
)

func TestFetchWithTimeout(t *testing.T) {
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
	result := fetchWithTimeout(urls, 300*time.Millisecond)
	fmt.Println(result)
	if k != len(result) {
		t.Error()
	}
}
