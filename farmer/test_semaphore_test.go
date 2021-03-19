package farmer

import (
	// TODO
	//
	"testing"
)

func TestSemaphore(t *testing.T) {
	semaphore1 := NewSemaphore(0)
	Assert(0 == semaphore1.AvailablePermits())

	semaphore1.Release()
	Assert(1 == semaphore1.AvailablePermits())

	semaphore1.ReleaseN(5)
	Assert(6 == semaphore1.AvailablePermits())

	semaphore1.ReleaseN(5)
	Assert(11 == semaphore1.AvailablePermits())

	Assert(semaphore1.Acquire())
	Assert(10 == semaphore1.AvailablePermits())

	Assert(semaphore1.AcquireN(5))
	Assert(5 == semaphore1.AvailablePermits())

	Assert(semaphore1.TryAcquire())
	Assert(4 == semaphore1.AvailablePermits())

	Assert(semaphore1.TryAcquireN(2))
	Assert(2 == semaphore1.AvailablePermits())

	Assert(!semaphore1.TryAcquireN(3))
	Assert(2 == semaphore1.AvailablePermits())

	// semaphore1.AcquireN(3)
	return
}
