package sequential

import "fmt"

type h2o struct {
	hReady chan struct{}
	oReady chan struct{}
	result chan struct{}
}

func newH2o() *h2o {
	h := &h2o{
		hReady: make(chan struct{}, 2),
		oReady: make(chan struct{}, 1),
		result: make(chan struct{}, 3),
	}
	go func() {
		for {
			// reset
			h.hReady <- struct{}{}
			h.hReady <- struct{}{}
			h.oReady <- struct{}{}
			// wait for h2o produced
			for i := 0; i < 3; i++ {
				// TODO: check closed check
				<-h.result
			}
		}
	}()
	return h
}

func (h2o *h2o) h() {
	<-h2o.hReady
	fmt.Print("H")
	h2o.result <- struct{}{}
}

func (h2o *h2o) o() {
	<-h2o.oReady
	fmt.Print("O")
	h2o.result <- struct{}{}
}
