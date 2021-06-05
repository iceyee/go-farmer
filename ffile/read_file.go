package ffile

import (
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/ftype"
	"io/ioutil"
	//
)

// 读文件.
func ReadFile(filename string) ([]byte, ftype.Error) {
	var content []byte
	var e error
	content, e = ioutil.ReadFile(filename)
	if nil != e {
		return nil, ferror.New(e)
	}
	return content, nil
}
