package futil

import (
	"sync"
	//
)

// 删除过期的键.
// @expiredTime - 过期时间, 时间戳, 单位秒, 小于这个时间的键被删除.
// @lock - 同步锁, 可以为nil.
func DeleteExpiredKey(
	data map[string]int64,
	expiredTime int64,
	lock *sync.RWMutex) {

	if nil != lock {
		lock.Lock()
		lock.Unlock()
	}
	var a001 []string
	a001 = make([]string, 0, len(data)+1)
	for key, value := range data {
		if value < expiredTime {
			a001 = append(a001, key)
		}
	}
	for _, value := range a001 {
		delete(data, value)
	}
	return
}

type T737 struct {
	Time  int64
	Value interface{}
}

// 功能同DeleteExpiredKey().
func DeleteExpiredKey2(
	data map[string]T737,
	expiredTime int64,
	lock *sync.RWMutex) {

	if nil != lock {
		lock.Lock()
		lock.Unlock()
	}
	var a001 []string
	a001 = make([]string, 0, len(data)+1)
	for key, value := range data {
		if value.Time < expiredTime {
			a001 = append(a001, key)
		}
	}
	for _, value := range a001 {
		delete(data, value)
	}
	return
}
