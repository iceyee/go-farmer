package fassert

import (
//
)

// 断言.
// 如果不通过, 则中断程序.
func Assert(condition bool, message ...string) {
	var a001 string
	for _, x := range message {
		a001 += x
		a001 += " "
	}
	if condition {
		println(a001 + " - 通过.")
	} else {
		println(a001 + " - 不通过.")
		panic(condition)
	}
	return
}

// 断言.
// 如果不通过, 也不会中断程序.
func Assert2(condition bool, message ...string) {
	var a001 string
	for _, x := range message {
		a001 += x
		a001 += " "
	}
	if condition {
		println(a001 + " - 通过.")
	} else {
		println(a001 + " - 不通过.")
	}
	return
}
