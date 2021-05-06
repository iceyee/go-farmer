package futil

import (
	"time"
	//
)

// 延时, 单位毫秒.
func Sleep(millisecond int64) {
	if millisecond <= 0 {
		millisecond = 1
	}
	time.Sleep(time.Duration(millisecond) * time.Millisecond)
	return
}
