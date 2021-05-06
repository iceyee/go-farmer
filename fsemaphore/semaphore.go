package fsemaphore

import (
//
)

// 信号量, 基于chan来实现功能.
type Semaphore struct {
	channel chan byte
}

func New(max int) *Semaphore {
	if max <= 0 {
		max = 1
	}
	var s *Semaphore
	s = new(Semaphore)
	s.channel = make(chan byte, max)
	return s
}

func (s *Semaphore) Acquire() bool {
	_, ok := <-s.channel
	return ok
}

func (s *Semaphore) AcquireN(n int) bool {
	for x := 0; x < n; x++ {
		_, ok := <-s.channel
		if !ok {
			return false
		}
	}
	return true
}

func (s *Semaphore) AvailablePermits() int {
	return len(s.channel)
}

func (s *Semaphore) Release() {
	s.channel <- 0
	return
}

func (s *Semaphore) ReleaseN(n int) {
	for x := 0; x < n; x++ {
		s.channel <- 0
	}
	return
}

func (s *Semaphore) TryAcquire() bool {
	if 0 == len(s.channel) {
		return false
	}
	_, ok := <-s.channel
	return ok
}

func (s *Semaphore) TryAcquireN(n int) bool {
	for x := 0; x < n; x++ {
		if len(s.channel) < n-x {
			return false
		}
		_, ok := <-s.channel
		if !ok {
			return false
		}
	}
	return true
}
