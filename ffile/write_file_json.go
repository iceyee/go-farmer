package ffile

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/ftype"
	"io/ioutil"
	//
)

// 写文件, 权限0664, json格式.
func WriteFileJson(
	filename string,
	data interface{}) ftype.Error {

	var content []byte
	var e error
	content, e = json.Marshal(data)
	if nil != e {
		return ferror.New(e)
	}
	e = ioutil.WriteFile(filename, content, 0664)
	if nil != e {
		return ferror.New(e)
	}
	return nil
}
