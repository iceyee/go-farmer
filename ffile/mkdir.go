package ffile

import (
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/ftype"
	"os"
	//
)

// 创建目录, 权限0775.
func Mkdir(path string) ftype.Error {
	var e error
	e = os.MkdirAll(path, 0775)
	if nil != e {
		e = ferror.New(e)
	}
	return e
}
