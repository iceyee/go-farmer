package fassert

import (
//
)

// 断言e == nil, 否则中断程序
func CheckError(e error) {
	if nil != e {
		panic(e)
	}
	return
}
