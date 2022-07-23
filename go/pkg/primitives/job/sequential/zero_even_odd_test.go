package sequential

import (
	"fmt"
	"testing"
)

func TestZeroEvenOdd(t *testing.T) {
	z := newZeroEvenOdd(1000)
	go func() {
		z.even()
	}()
	go func() {
		z.odd()
	}()
	go func() {
		z.zero()
	}()
	<-z.done()
	fmt.Println()
	fmt.Println("All printed!")
}
