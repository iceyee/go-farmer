package fweb

import (
//
)

// 响应模式.
type ResponseMode int64

const (
	M_JSON = iota
	M_PLAIN
)

var mode ResponseMode = M_PLAIN

// 设置响应模式, M_JSON表示返回的是json格式.
func SetResponseMode(m ResponseMode) {
	if M_JSON != m &&
		M_PLAIN != m {

		panic(0)
	}
	mode = m
	return
}
