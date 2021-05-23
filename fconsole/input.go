package fconsole

import (
	"github.com/iceyee/go-farmer/v4/ferror"
	"github.com/iceyee/go-farmer/v4/ftype"
	"os"
	"strings"
	//
)

// 从stdin获得输入.
// @destination - 接收输入的变量.
// @message - 提示输入.
func Input(destination *string, message string) ftype.Error {
	if "" != message {
		println(message)
	}
	print(">>> ")
	var a001 []byte
	a001 = make([]byte, 0xffff)
	var a002 int
	var e error
	a002, e = os.Stdin.Read(a001)
	if nil != e {
		return ferror.New(e)
	}
	var a003 string
	a003 = strings.Trim(string(a001[0:a002]), " \r\n\t\v")
	if nil != destination {
		*destination = a003
	}
	return nil
}
