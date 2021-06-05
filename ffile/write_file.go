package ffile

import (
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/ftype"
	"io/ioutil"
	//
)

// 写文件, 权限0664.
func WriteFile(filename string, data []byte) ftype.Error {
	var e error
	e = ioutil.WriteFile(filename, data, 0664)
	if nil != e {
		return ferror.New(e)
	}
	return nil
}
