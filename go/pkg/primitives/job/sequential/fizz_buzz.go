package sequential

import "fmt"

type FizzBuzz interface {
	Fizz()
	Buzz()
	FizzBuzz()
	Number()
}

type fb struct {
	n    int
	fCh  chan int
	bCh  chan int
	fbCh chan int
	nCh  chan int
}

func (f *fb) Fizz() {
	for {
		v, ok := <-f.fCh
		if !ok {
			return
		}
		fmt.Println("fizz")
		f.dispatchNext(v)
	}
}

func (f *fb) Buzz() {
	for {
		v, ok := <-f.bCh
		if !ok {
			return
		}
		fmt.Println("buzz")
		f.dispatchNext(v)
	}
}

func (f *fb) FizzBuzz() {
	for {
		v, ok := <-f.fbCh
		if !ok {
			return
		}
		fmt.Println("fizzbuzz")
		f.dispatchNext(v)
	}
}

func (f *fb) Number() {
	for {
		v, ok := <-f.nCh
		if !ok {
			return
		}
		fmt.Println(v)
		f.dispatchNext(v)
	}
}

func (f *fb) dispatchNext(v int) {
	if v == f.n {
		f.closeAll()
		return
	}
	v++
	if v%15 == 0 {
		f.fbCh <- v
	} else if v%5 == 0 {
		f.bCh <- v
	} else if v%3 == 0 {
		f.fCh <- v
	} else {
		f.nCh <- v
	}
}

func (f *fb) init() {
	f.nCh <- 1
}

func (f *fb) closeAll() {
	close(f.fCh)
	close(f.bCh)
	close(f.fbCh)
	close(f.nCh)
}

func NewFizzBuzz(n int) FizzBuzz {
	fizzBuzz := &fb{
		n:    n,
		fCh:  make(chan int, 1),
		bCh:  make(chan int, 1),
		fbCh: make(chan int, 1),
		nCh:  make(chan int, 1),
	}
	fizzBuzz.init()
	return fizzBuzz
}
