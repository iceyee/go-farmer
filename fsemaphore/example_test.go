package fsemaphore

import (
	"github.com/iceyee/go-farmer/v5/fassert"
	"testing"
	//
)

func TestSemaphore(t *testing.T) {
	var s *Semaphore
	s = New(0xfff)
	fassert.Assert(0 == s.AvailablePermits(), "初始0.")
	s.Release()
	fassert.Assert(1 == s.AvailablePermits(), "释放, +1, =1.")
	s.ReleaseN(5)
	fassert.Assert(6 == s.AvailablePermits(), "释放, +5, =6.")
	s.ReleaseN(5)
	fassert.Assert(11 == s.AvailablePermits(), "释放, +5, =11.")
	s.Acquire()
	fassert.Assert(10 == s.AvailablePermits(), "获取, -1, =10.")
	s.AcquireN(5)
	fassert.Assert(5 == s.AvailablePermits(), "获取, -5, =5.")
	fassert.Assert(s.TryAcquire(), "成功尝试获取1个.")
	fassert.Assert(4 == s.AvailablePermits(), "获取后, -1, =4.")
	return
}

func Example() {
	var s *Semaphore
	s = New(0xfff)
	fassert.Assert(0 == s.AvailablePermits(), "初始0.")
	s.Release()
	fassert.Assert(1 == s.AvailablePermits(), "释放, +1, =1.")
	s.ReleaseN(5)
	fassert.Assert(6 == s.AvailablePermits(), "释放, +5, =6.")
	s.ReleaseN(5)
	fassert.Assert(11 == s.AvailablePermits(), "释放, +5, =11.")
	s.Acquire()
	fassert.Assert(10 == s.AvailablePermits(), "获取, -1, =10.")
	s.AcquireN(5)
	fassert.Assert(5 == s.AvailablePermits(), "获取, -5, =5.")
	fassert.Assert(s.TryAcquire(), "成功尝试获取1个.")
	fassert.Assert(4 == s.AvailablePermits(), "获取后, -1, =4.")
	return
}
