package sequential

import (
	"fmt"
	"sync/atomic"
)

type zeroEvenOdd struct {
	nextOdd  uint64
	nextEven uint64
	n        int
	fin      chan struct{}
}

func newZeroEvenOdd(n int) *zeroEvenOdd {
	return &zeroEvenOdd{
		nextOdd:  1,
		nextEven: 2,
		n:        n,
		fin:      make(chan struct{}),
	}
}

func (z *zeroEvenOdd) zero() {
	fmt.Print(0)
}
func (z *zeroEvenOdd) even() {
	fmt.Print(z.nextEven)
	atomic.AddUint64(&z.nextEven, 2)
}

func (z *zeroEvenOdd) odd() {
	fmt.Print(z.nextOdd)
	atomic.AddUint64(&z.nextOdd, 2)
}

func (z *zeroEvenOdd) done() <-chan struct{} {
	return z.fin
}
