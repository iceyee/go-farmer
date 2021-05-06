package fassert

import (
	"testing"
	//
)

func ExampleAssert(t *testing.T) {
	Assert(true, "用 Assert() 测试true")
	func() {
		defer func() {
			if nil == recover() {
				t.Fatal("Assert(false) 有问题")
			}
		}()
		Assert(false, "用 Assert() 测试false")
	}()
	return
}

func ExampleAssert2(t *testing.T) {
	Assert2(true, "用 Assert2() 测试true")
	Assert2(false, "用 Assert2() 测试false")
	return
}
