package flog

import (
	"github.com/iceyee/go-farmer/v4/ftype"
	"strings"
	//
)

func process(message interface{}) string {
	var result string
	switch message.(type) {

	case string:
		result = message.(string)
	case error:
		result = message.(error).Error()
	case ftype.Stringer:
		result = message.(ftype.Stringer).String()
	case []byte:
		result = string(message.([]byte))
	default:
		result = "[Unkown]"
	}
	result = strings.Replace(result, "\n", "\n    ", -1)
	return result
}

func Debug(message interface{}) {
	if DEBUG < logLevel {
		return
	}
	var a001 string
	a001 = (*time355) + "  DEBUG  #  " + process(message) + "\n"
	debugWriter.Write([]byte(a001))
	return
}

// message类型支持, string, error, ftype.Stringer, []byte.
func Info(message interface{}) {
	if INFO < logLevel {
		return
	}
	var a001 string
	a001 = (*time355) + "  INFO   #  " + process(message) + "\n"
	infoWriter.Write([]byte(a001))
	return
}

func Warn(message interface{}) {
	if WARN < logLevel {
		return
	}
	var a001 string
	a001 = (*time355) + "  WARN   #  " + process(message) + "\n"
	warnWriter.Write([]byte(a001))
	return
}

func Error(message interface{}) {
	if ERROR < logLevel {
		return
	}
	var a001 string
	a001 = (*time355) + "  ERROR  #  " + process(message) + "\n"
	errorWriter.Write([]byte(a001))
	return
}

func Fatal(message interface{}) {
	if FATAL < logLevel {
		return
	}
	var a001 string
	a001 = (*time355) + "  FATAL  #  " + process(message) + "\n"
	fatalWriter.Write([]byte(a001))
	return
}
