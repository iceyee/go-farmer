package flog

import (
	"github.com/iceyee/go-farmer/v5/ffile"
	"github.com/iceyee/go-farmer/v5/fschedule"
	"os"
	"time"
	//
)

func init() {
	fschedule.Schedule(
		"1 0 * * *",
		1000,
		true,
		deleteOldFile)
	return
}

var x393 int64 = 8

// 删除几天前的文件, 默认8天.
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
		} else if 1*60*60*24*x393 < time.Now().Unix()-file.ModTime().Unix() {
			var a002 string
			a002 = ffile.Path(
				ffile.HomeDirectory,
				"farmer",
				"share",
				"farmer-log",
				projectName,
				file.Name())
			os.Remove(a002)
		}
	}
	return
}

// 设定, 要保留几天, 默认8天.
func SetSaveDays(a int64) {
	x393 = a
	return
}
