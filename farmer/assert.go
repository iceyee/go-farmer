package farmer

import (
	"log"
	"runtime/debug"
	//
)

func Assert(condition bool) {
	if !condition {
		var message = "断言失败\n" + string(debug.Stack())
		log.Fatalln(message)
	}
	return
}

func CheckError(e error) {
	if nil != e {
		log.Fatalln(e)
	}
	return
}
