package fschedule

import (
	"github.com/iceyee/go-farmer/v4/futil"
	"testing"
	//
)

func TestSchedule(t *testing.T) {
	Schedule("0-59/1 0-23 * * *", 200, false, func() {
		println("hello world")
	})
	futil.Sleep(1000)
	return
}

func ExampleSchedule() {
	Schedule("0-59/1 0-23 * * *", 200, false, func() {
		println("hello world")
	})
	futil.Sleep(1000)
	return
}
