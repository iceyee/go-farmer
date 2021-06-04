package flog

import (
	"testing"
	//
)

func TestLog(t *testing.T) {
	Debug("hello world.")
	Info("hello world.")
	Warn("hello world.")
	Error("hello world.")
	Fatal("hello world.")

	SetLogLevel(L_INFO)
	SetProjectName("GO-FARMER", F_WARN|F_FATAL)
	Debug("hello world.")
	Info("hello world.")
	Warn("hello world.")
	Error("hello world.")
	Fatal("hello world.")
	return
}

func Example() {
	Debug("hello world.")
	Info("hello world.")
	Warn("hello world.")
	Error("hello world.")
	Fatal("hello world.")

	SetLogLevel(L_INFO)
	SetProjectName("GO-FARMER", F_WARN|F_FATAL)
	Debug("hello world.")
	Info("hello world.")
	Warn("hello world.")
	Error("hello world.")
	Fatal("hello world.")
	return
}
