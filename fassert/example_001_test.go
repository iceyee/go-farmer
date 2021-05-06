package fassert

import (
	"errors"
	//
)

func ExampleAssert() {
	Assert(true, "用 Assert() 测试true")
	func() {
		defer func() {
			if nil == recover() {
				println("Assert(false) 有问题")
			}
		}()
		Assert(false, "用 Assert() 测试false")
	}()
	return
}

func ExampleAssert2() {
	Assert2(true, "用 Assert2() 测试true")
	Assert2(false, "用 Assert2() 测试false")
	return
}

func ExampleCheckError() {
	func() {
		defer func() {
			if e := recover(); nil == e {
				println("CheckError() 有问题")
			} else {
				println(e.(error).Error())
			}
		}()
		CheckError(errors.New("hello world."))
	}()
	return
}
