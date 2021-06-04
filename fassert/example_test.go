package fassert

import (
	"errors"
	"testing"
	//
)

func Test(t *testing.T) {
	return
}

func TestAssert(t *testing.T) {
	Assert2(true, "用 Assert2() 测试true.")
	Assert2(false, "用 Assert2() 测试false.")
	Assert(true, "用 Assert() 测试true.")
	func() {
		defer func() {
			if nil != recover() {
				t.Fatal("Assert() 如同预期, 出现异常.")
			} else {
				t.Fatal("Assert() 非预期, 没出现异常.")
			}
		}()
		Assert(false, "用 Assert() 测试false.")
	}()
	return
}

func TestCheckError(t *testing.T) {
	func() {
		defer func() {
			if e := recover(); nil != e {
				t.Log(e)
			} else {
				t.Fatal("CheckError() 非预期中断.")
			}
		}()
		CheckError(errors.New("hello world."), "测试异常.")
	}()
	return
}

func ExampleAssert() {
	Assert(true, "用 Assert() 测试true.")
	func() {
		defer func() {
			if nil != recover() {
				println("Assert() 如同预期, 出现异常.")
			} else {
				println("Assert() 非预期, 没出现异常.")
			}
		}()
		Assert(false, "用 Assert() 测试false.")
	}()
	return
}

func ExampleAssert2() {
	Assert2(true, "用 Assert2() 测试true.")
	Assert2(false, "用 Assert2() 测试false.")
	return
}

func ExampleCheckError() {
	func() {
		defer func() {
			if e := recover(); nil != e {
				println(e.(error).Error())
			} else {
				println("CheckError() 非预期中断.")
			}
		}()
		CheckError(errors.New("hello world."), "测试异常.")
	}()
	return
}
