package fassert

import (
	"github.com/iceyee/go-farmer/v3/flog"
	//
)

// 断言.
// 如果不通过, 则中断程序.
func Assert(condition bool, message string) {
	if condition {
		flog.Debug(message + " - 通过.")
	} else {
		flog.Debug(message + " - 不通过.")
		panic(condition)
	}
	return
}

// 断言.
// 如果不通过, 也不会中断程序.
func Assert2(condition bool, message string) {
	if condition {
		flog.Debug(message + " - 通过.")
	} else {
		flog.Debug(message + " - 不通过.")
	}
	return
}
