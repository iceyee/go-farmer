package fsemaphore

import (
	"github.com/iceyee/go-farmer/v3/fassert"
	"testing"
	//
)

func TestSemaphore(t *testing.T) {
	var s *Semaphore
	s = New(0xfff)
	fassert.Assert(0 == s.AvailablePermits(), "")
	s.Release()
	fassert.Assert(1 == s.AvailablePermits(), "")
	s.ReleaseN(5)
	fassert.Assert(6 == s.AvailablePermits(), "")
	s.ReleaseN(5)
	fassert.Assert(11 == s.AvailablePermits(), "")
	fassert.Assert(s.Acquire(), "")
	fassert.Assert(10 == s.AvailablePermits(), "")
	fassert.Assert(s.AcquireN(5), "")
	fassert.Assert(5 == s.AvailablePermits(), "")
	fassert.Assert(s.TryAcquire(), "")
	fassert.Assert(4 == s.AvailablePermits(), "")
	fassert.Assert(s.TryAcquireN(2), "")
	fassert.Assert(2 == s.AvailablePermits(), "")
	fassert.Assert(!s.TryAcquireN(3), "")
	fassert.Assert(2 == s.AvailablePermits(), "")
	// s.AcquireN(3)
	return
}

func Example() {
	var s *Semaphore
	s = New(0xfff)
	fassert.Assert(0 == s.AvailablePermits(), "")
	s.Release()
	fassert.Assert(1 == s.AvailablePermits(), "")
	s.ReleaseN(5)
	fassert.Assert(6 == s.AvailablePermits(), "")
	s.ReleaseN(5)
	fassert.Assert(11 == s.AvailablePermits(), "")
	fassert.Assert(s.Acquire(), "")
	fassert.Assert(10 == s.AvailablePermits(), "")
	fassert.Assert(s.AcquireN(5), "")
	fassert.Assert(5 == s.AvailablePermits(), "")
	fassert.Assert(s.TryAcquire(), "")
	fassert.Assert(4 == s.AvailablePermits(), "")
	fassert.Assert(s.TryAcquireN(2), "")
	fassert.Assert(2 == s.AvailablePermits(), "")
	fassert.Assert(!s.TryAcquireN(3), "")
	fassert.Assert(2 == s.AvailablePermits(), "")
	// s.AcquireN(3)
	return
}
