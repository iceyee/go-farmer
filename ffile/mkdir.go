package ffile

import (
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/ftype"
	"os"
	//
)

// 创建目录
func Mkdir(path string) ftype.Error {
	var e error
	e = os.MkdirAll(path, 0775)
	if nil != e {
		e = ferror.New(e)
	}
	return e
}
