package ftimer

import (
	"time"
	//
)

// 计时器
type Timer struct {
	begin *int64
}

// 生成计时器, 并开始计时.
func New() Timer {
	var t Timer
	t.begin = new(int64)
	*t.begin = time.Now().UnixNano() / 1000000
	return t
}

// 重置计时.
func (t Timer) Reset() {
	*t.begin = time.Now().UnixNano() / 1000000
	return
}

// 计时.
func (t Timer) Timing() int64 {
	return time.Now().UnixNano()/1000000 - (*t.begin)
}
