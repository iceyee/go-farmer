package ffile

import (
	"github.com/iceyee/go-farmer/v3/fassert"
	//
)

func ExampleWriteFile() {
	e := WriteFile("/tmp/go-farmer-test.txt", []byte("hello world."))
	fassert.CheckError(e)
	return
}

func ExampleReadFile() {
	content, e := ReadFile("/tmp/go-farmer-test.txt")
	fassert.CheckError(e)
	println(string(content))
	return
}

func ExamplePath() {
	println(Path(HomeDirectory, "go-farmer", "test.txt"))
	return
}

func ExampleInstallDirectory() {
	InstallDirectory()
	return
}
