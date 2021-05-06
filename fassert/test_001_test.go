package fassert

import (
	"errors"
	"testing"
	//
)

func TestAssert(t *testing.T) {
	Assert2(true, "用 Assert2() 测试true")
	Assert2(false, "用 Assert2() 测试false")
	Assert(true, "用 Assert() 测试true")
	func() {
		defer func() {
			if nil == recover() {
				t.Fatal("Assert() 有问题")
			}
		}()
		Assert(false, "用 Assert() 测试false")
	}()
	return
}

func TestCheckError(t *testing.T) {
	func() {
		defer func() {
			if e := recover(); nil == e {
				t.Fatal("CheckError() 有问题")
			} else {
				t.Log(e)
			}
		}()
		CheckError(errors.New("hello world."))
	}()
	return
}
