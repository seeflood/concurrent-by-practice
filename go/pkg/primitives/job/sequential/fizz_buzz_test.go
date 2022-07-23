package sequential

import (
	"fmt"
	"sync"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	fb := NewFizzBuzz(100)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		fb.FizzBuzz()
		wg.Done()
	}()
	go func() {
		fb.Fizz()
		wg.Done()
	}()
	go func() {
		fb.Buzz()
		wg.Done()
	}()
	go func() {
		fb.Number()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Finish!")
}
