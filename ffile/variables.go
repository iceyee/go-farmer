package ffile

import (
	"os"
	"os/user"
	//
)

// 用户主目录.
var HomeDirectory string

// 路径分隔符.
var PathSeparator string = string(os.PathSeparator)

// 多路径分隔符.
// 比如, linux下是':'.
var PathListSeparator string = string(os.PathListSeparator)

func init() {
	var user001 *user.User
	var e error
	user001, e = user.Current()
	if nil != e {
		panic(e)
	}
	HomeDirectory = user001.HomeDir
	return
}
