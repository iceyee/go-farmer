package farmer

import (
//
)

// expiredTime - 过期时间, 时间戳, 单位秒, 小于这个时间的键被删除
func DeleteExpiredKey(data map[string]int64, expiredTime int64) {
	var a1 []string
	a1 = make([]string, 0, len(data))
	for key, value := range data {
		if value < expiredTime {
			a1 = append(a1, key)
		}
	}
	for index, value := range a1 {
		delete(data, value)
		_ = index
	}
	return
}
