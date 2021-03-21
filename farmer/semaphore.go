package farmer

import (
// TODO
//
)

// 信号量, 但是不能超过0xff, 基于chan来实现功能
type Semaphore struct {
	channel chan byte
}

func NewSemaphore(permits int) *Semaphore {
	s := new(Semaphore)
	s.channel = make(chan byte, 0xff)
	for x := 0; x < permits; x++ {
		s.channel <- 0
	}
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
