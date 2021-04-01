package farmer

import (
	"testing"
	"time"
	//
)

func TestExpiredKey(t *testing.T) {
	var a1 map[string]int64
	a1 = make(map[string]int64, 0xff)
	a1["a"] = 0
	a1["b"] = time.Now().Unix() - 10
	a1["c"] = time.Now().Unix() - 5
	a1["d"] = time.Now().Unix()
	t.Log(a1)
	Assert(4 == len(a1))
	DeleteExpiredKey(a1, time.Now().Unix()-6)
	t.Log(a1)
	Assert(2 == len(a1))
	return
}
