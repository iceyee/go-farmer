package ffile

import (
	"github.com/iceyee/go-farmer/v3/fassert"
	"testing"
	//
)

func TestWriteFile(t *testing.T) {
	e := WriteFile("/tmp/go-farmer-test.txt", []byte("hello world."))
	fassert.CheckError(e)
	return
}

func TestReadFile(t *testing.T) {
	content, e := ReadFile("/tmp/go-farmer-test.txt")
	fassert.CheckError(e)
	t.Log(string(content))
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
