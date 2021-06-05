package ffile

import (
	"github.com/iceyee/go-farmer/v5/fassert"
	"testing"
	//
)

func Test(t *testing.T) {
	return
}

func TestWriteFile(t *testing.T) {
	e := WriteFile("/tmp/go-farmer-test.txt", []byte("hello world."))
	fassert.CheckError(e, "写文件/tmp/go-farmer-test.txt")
	return
}

func TestReadFile(t *testing.T) {
	content, e := ReadFile("/tmp/go-farmer-test.txt")
	fassert.CheckError(e, "读文件/tmp/go-farmer-test.txt")
	t.Log(string(content))
	return
}

func TestWriteFileJson(t *testing.T) {
	type T struct {
		A string
		B int64
	}
	var a T
	a.A = "hello world"
	a.B = 10086
	e := WriteFileJson("/tmp/go-farmer-test-json.txt", a)
	fassert.CheckError(e, "写文件/tmp/go-farmer-test-json.txt")
	return
}

func TestReadFileJson(t *testing.T) {
	type T struct {
		A string
		B int64
	}
	var a T
	e := ReadFileJson("/tmp/go-farmer-test-json.txt", &a)
	fassert.CheckError(e, "读文件/tmp/go-farmer-test-json.txt")
	t.Log(a.A)
	t.Log(a.B)
	return
}

func TestPath(t *testing.T) {
	t.Log(Path(HomeDirectory, "go-farmer", "test.txt"))
	return
}

func TestInstallDirectory(t *testing.T) {
	InstallDirectory()
	return
}

func ExampleWriteFile() {
	e := WriteFile("/tmp/go-farmer-test.txt", []byte("hello world."))
	fassert.CheckError(e, "写文件/tmp/go-farmer-test.txt")
	return
}

func ExampleReadFile() {
	content, e := ReadFile("/tmp/go-farmer-test.txt")
	fassert.CheckError(e, "读文件/tmp/go-farmer-test.txt")
	println(string(content))
	return
}

func ExampleWriteFileJson() {
	type T struct {
		A string
		B int64
	}
	var a T
	a.A = "hello world"
	a.B = 10086
	e := WriteFileJson("/tmp/go-farmer-test-json.txt", a)
	fassert.CheckError(e, "写文件/tmp/go-farmer-test-json.txt")
	return
}

func ExampleReadFileJson() {
	type T struct {
		A string
		B int64
	}
	var a T
	e := ReadFileJson("/tmp/go-farmer-test-json.txt", &a)
	fassert.CheckError(e, "读文件/tmp/go-farmer-test-json.txt")
	println(a.A)
	println(a.B)
	return
}

func ExampleInstallDirectory() {
	InstallDirectory()
	return
}

func ExampleMkdir() {
	Mkdir("/tmp/tttt")
	return
}

func ExamplePath() {
	println(Path(HomeDirectory, "go-farmer", "test.txt"))
	return
}
