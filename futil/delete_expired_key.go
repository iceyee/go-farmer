package futil

import (
	"sync"
	"time"
	//
)

// 删除过期的键.
// expiredTime - 过期时间, 时间戳, 单位秒, 小于这个时间的键被删除.
func DeleteExpiredKey(data map[string]int64, expiredTime int64) {
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
func DeleteExpiredKey2(data map[string]T737, expiredTime int64) {
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

// 功能同DeleteExpiredKey(), value类型必须是int64.
// 后台线程循环执行, 不用重复调用.
func DeleteExpiredKey3(data *sync.Map, expiredTime int64) {
	go func() {
		for true {
			time.Sleep(1 * time.Minute)
			var a001 []interface{}
			a001 = make([]interface{}, 0, 0xfff)
			data.Range(func(key, value interface{}) bool {
				if value.(int64) < expiredTime {
					a001 = append(a001, key)
				}
				return true
			})
			for _, x := range a001 {
				data.Delete(x)
			}
		}
	}()
	return
}

// 功能同DeleteExpiredKey(), value类型必须是T737.
// 后台线程循环执行, 不用重复调用.
func DeleteExpiredKey4(data *sync.Map, expiredTime int64) {
	go func() {
		for true {
			time.Sleep(1 * time.Minute)
			var a001 []interface{}
			a001 = make([]interface{}, 0, 0xfff)
			data.Range(func(key, value interface{}) bool {
				if value.(T737).Time < expiredTime {
					a001 = append(a001, key)
				}
				return true
			})
			for _, x := range a001 {
				data.Delete(x)
			}
		}
	}()
	return
}
