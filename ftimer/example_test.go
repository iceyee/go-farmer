package ftimer

import (
	"github.com/iceyee/go-farmer/v4/fassert"
	"github.com/iceyee/go-farmer/v4/futil"
	"testing"
	//
)

func TestTimer(t *testing.T) {
	var timer Timer
	timer = New()
	futil.Sleep(1000)
	var a int64
	a = timer.Timing()
	fassert.Assert(999 <= a && a <= 1001, "计时1秒")
	timer.Reset()
	futil.Sleep(2000)
	a = timer.Timing()
	fassert.Assert(1999 <= a && a <= 2001, "计时2秒")
	return
}

func ExampleTimer() {
	var timer Timer
	timer = New()
	futil.Sleep(1000)
	var a int64
	a = timer.Timing()
	fassert.Assert(999 <= a && a <= 1001, "计时1秒")
	timer.Reset()
	futil.Sleep(2000)
	a = timer.Timing()
	fassert.Assert(1999 <= a && a <= 2001, "计时2秒")
	return
}
