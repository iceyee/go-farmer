package ffile

import (
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/ftype"
	"io/ioutil"
	//
)

// 读文件
func ReadFile(filename string) ([]byte, ftype.Error) {
	var content []byte
	var e ftype.Error
	content, e = ioutil.ReadFile(filename)
	if nil != e {
		e = ferror.New(e)
	}
	return content, e
}
