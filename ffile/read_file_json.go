package ffile

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/ftype"
	"io/ioutil"
	//
)

// 读文件, json格式.
// @data - 同json.Unmarshal()的参数.
func ReadFileJson(
	filename string,
	data interface{}) ftype.Error {

	var content []byte
	var e ftype.Error
	content, e = ioutil.ReadFile(filename)
	if nil != e {
		return ferror.New(e)
	}
	e = json.Unmarshal(content, data)
	if nil != e {
		return ferror.New(e)
	}
	return nil
}
