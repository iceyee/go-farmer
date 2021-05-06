package flog

import (
	"io"
	"os"
	//
)

// 日志等级
type LogLevel int64

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
	NONE
)

type LogFlag int64

const (
	F_DEBUG LogFlag = 1 << iota
	F_INFO
	F_WARN
	F_ERROR
	F_FATAL
)

var logFlag LogFlag
var logLevel LogLevel = DEBUG
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