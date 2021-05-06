package flog

import (
//
)

func ExampleLogLevel() {
	Debug("hello world.")
	Info("hello world.")
	Warn("hello world.")
	Error("hello world.")
	Fatal("hello world.")

	SetLogLevel(INFO)
	SetProjectName("TEST", F_WARN|F_FATAL)
	Debug("hello world.")
	Info("hello world.")
	Warn("hello world.")
	Error("hello world.")
	Fatal("hello world.")
	return
}
