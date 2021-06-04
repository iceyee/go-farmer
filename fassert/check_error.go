package fassert

import (
//
)

// 断言e == nil, 否则中断程序
func CheckError(e error, message ...string) {
	var a001 string
	for _, x := range message {
		a001 += x
		a001 += " "
	}
	if nil != e {
		println(a001 + " - 出现异常")
		panic(e)
	} else {
		println(a001 + " - 没有异常")
	}
	return
}
