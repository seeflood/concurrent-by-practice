package lock

type Lock struct {
	ch chan struct{}
}

func NewLock() *Lock {
	return &Lock{ch: make(chan struct{}, 1)}
}

func (l *Lock) TryLock() bool {
	select {
	case l.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

func (l *Lock) Lock() {
	l.ch <- struct{}{}
}

func (l *Lock) Unlock() {
	<-l.ch
}
