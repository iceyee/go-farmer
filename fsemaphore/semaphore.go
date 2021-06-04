package fsemaphore

import (
//
)

// 信号量, 基于chan来实现功能.
type Semaphore struct {
	channel chan byte
}

func New(max int64) *Semaphore {
	if max <= 0 {
		max = 1
	}
	var s *Semaphore
	s = new(Semaphore)
	s.channel = make(chan byte, max)
	return s
}

func (s *Semaphore) Acquire() {
	<-s.channel
	return
}

func (s *Semaphore) AcquireN(n int64) {
	for x := int64(0); x < n; x++ {
		<-s.channel
	}
	return
}

func (s *Semaphore) AvailablePermits() int64 {
	return int64(len(s.channel))
}

func (s *Semaphore) Release() {
	s.channel <- 0
	return
}

func (s *Semaphore) ReleaseN(n int64) {
	for x := int64(0); x < n; x++ {
		s.channel <- 0
	}
	return
}

func (s *Semaphore) TryAcquire() bool {
	if 0 == len(s.channel) {
		return false
	}
	<-s.channel
	return true
}
