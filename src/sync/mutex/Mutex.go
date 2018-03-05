package mutex

import "sync/semaphore"

type Mutex interface {
	Lock()
	Unlock()
}

type mutex struct {
	lock semaphore.Semaphore
}

func New() Mutex {
	// Ignore error.
	sem, _ := semaphore.New(1)
	return &mutex{
		lock: sem,
	}
}

func (mu *mutex) Lock() {
	mu.lock.Wait(1)
}

func (mu *mutex) Unlock() {
	mu.lock.Signal(1)
}
