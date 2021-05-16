package flog

import (
	"github.com/iceyee/go-farmer/v4/ffile"
	"github.com/iceyee/go-farmer/v4/fschedule"
	"os"
	"time"
	//
)

func init() {
	fschedule.Schedule(
		"1 0 * * *",
		1000,
		true,
		deleteOldFile,
	)
	return
}

// 删除7-8天前的文件.
func deleteOldFile() {
	if "" == projectName {
		return
	}
	var a001 string
	a001 = ffile.Path(
		ffile.HomeDirectory,
		"farmer",
		"share",
		"farmer-log",
		projectName)
	var directory001 *os.File
	var e error
	directory001, e = os.Open(a001)
	if nil != e {
		panic(e)
	}
	var state001 os.FileInfo
	state001, e = directory001.Stat()
	if nil != e {
		panic(e)
	} else if !state001.IsDir() {
		panic(a001 + "不是目录.")
	}
	var files []os.FileInfo
	files, e = directory001.Readdir(0)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else if 1*60*60*24*8 < time.Now().Unix()-file.ModTime().Unix() {
			os.Remove(file.Name())
		}
	}
	return
}
