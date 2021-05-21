package futil

import (
	"github.com/iceyee/go-farmer/v4/ferror"
	"github.com/iceyee/go-farmer/v4/ftype"
	"os"
	"strings"
	//
)

// 从stdin获得输入.
func Input() (string, ftype.Error) {
	var a001 []byte
	a001 = make([]byte, 0xffff)
	var a002 int
	var e error
	a002, e = os.Stdin.Read(a001)
	if nil != e {
		return "", ferror.New(e)
	}
	return strings.Trim(string(a001[0:a002]), " \r\n\t\v"), nil
}
