package flog

import (
	"github.com/iceyee/go-farmer/v3/fschedule"
	"os"
	//
)

func init() {
	fschedule.Schedule(
		"* * * * *",
		1000,
		true,
		synchronize,
	)
	return
}

// 同步文件.
func synchronize() {
	for _, file := range []*os.File{
		debugFile, infoFile, warnFile, errorFile, fatalFile} {

		if nil != file {
			file.Sync()
		}
	}
	return
}
