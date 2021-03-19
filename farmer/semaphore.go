package farmer

import (
// TODO
//
)

// 信号量, 但是不能超过0xff
type Semaphore struct {
	channel chan byte
}

func NewSemaphore(permits int) *Semaphore {
	semaphore1 := new(Semaphore)
	semaphore1.channel = make(chan byte, 0xff)
	for x := 0; x < permits; x++ {
		semaphore1.channel <- 0
	}
	return semaphore1
}

func (semaphore *Semaphore) Acquire() bool {
	_, ok := <-semaphore.channel
	return ok
}

func (semaphore *Semaphore) AcquireN(n int) bool {
	for x := 0; x < n; x++ {
		_, ok := <-semaphore.channel
		if !ok {
			return false
		}
	}
	return true
}

func (semaphore *Semaphore) AvailablePermits() int {
	return len(semaphore.channel)
}

func (semaphore *Semaphore) Release() {
	semaphore.channel <- 0
	return
}

func (semaphore *Semaphore) ReleaseN(n int) {
	for x := 0; x < n; x++ {
		semaphore.channel <- 0
	}
	return
}

func (semaphore *Semaphore) TryAcquire() bool {
	if 0 == len(semaphore.channel) {
		return false
	}
	_, ok := <-semaphore.channel
	return ok
}

func (semaphore *Semaphore) TryAcquireN(n int) bool {
	for x := 0; x < n; x++ {
		if len(semaphore.channel) < n-x {
			return false
		}
		_, ok := <-semaphore.channel
		if !ok {
			return false
		}
	}
	return true
}
