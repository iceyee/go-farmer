package flog

import (
	"io"
	"os"
	//
)

// 日志等级. 低于这个等级的不会输出.
type LogLevel int64

const (
	L_DEBUG LogLevel = iota
	L_INFO
	L_WARN
	L_ERROR
	L_FATAL
	L_NONE
)

// 日志标志. 表示哪些日志需要记录.
type LogFlag int64

const (
	F_DEBUG LogFlag = 1 << iota
	F_INFO
	F_WARN
	F_ERROR
	F_FATAL
)

var logFlag LogFlag
var logLevel LogLevel = L_DEBUG
var projectName string
var time355 *string = new(string)

var debugWriter io.Writer = os.Stdout
var infoWriter io.Writer = os.Stdout
var warnWriter io.Writer = os.Stdout
var errorWriter io.Writer = os.Stderr
var fatalWriter io.Writer = os.Stderr

var debugFile *os.File
var infoFile *os.File
var warnFile *os.File
var errorFile *os.File
var fatalFile *os.File
