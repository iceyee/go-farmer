package flog

import (
//
)

// 设置日志等级. 低于这个等级的日志不会被输出.
func SetLogLevel(level LogLevel) {
	logLevel = level
	return
}
