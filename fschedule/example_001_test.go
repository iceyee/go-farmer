package fschedule

import (
	"github.com/iceyee/go-farmer/v3/futil"
	//
)

func ExampleSchedule() {
	Schedule("0-59/1 0-23 * * *", 200, false, func() {
		println("hello world")
	})
	futil.Sleep(1000)
	return
}
