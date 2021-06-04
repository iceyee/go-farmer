package futil

import (
	"time"
	//
)

// 当前时间, 单位:秒, 1秒刷新10次.
var NowTime int64 = time.Now().Unix()

// 当前时间, 单位:毫秒, 1秒刷新10次.
var NowTime2 int64 = time.Now().UnixNano() / 1000000

func init() {
	go func() {
		for true {
			time.Sleep(100 * time.Millisecond)
			NowTime2 = time.Now().UnixNano() / 1000000
			NowTime = NowTime2 / 1000
		}
	}()
	return
}
