package futil

import (
	"time"
	//
)

// 当前时间, 1秒刷新10次.
var NowTime int64 = time.Now().Unix()

func init() {
	go func() {
		for true {
			time.Sleep(100 * time.Millisecond)
			NowTime = time.Now().Unix()
		}
	}()
	return
}
