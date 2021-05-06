package ffile

import (
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/ftype"
	"io/ioutil"
	//
)

// 写文件
func WriteFile(filename string, data []byte) ftype.Error {
	var e ftype.Error
	e = ioutil.WriteFile(filename, data, 0664)
	if nil != e {
		e = ferror.New(e)
	}
	return e
}
