package flog

import (
	"fmt"
	"github.com/iceyee/go-farmer/v3/ffile"
	"github.com/iceyee/go-farmer/v3/fschedule"
	"io"
	"os"
	"time"
	//
)

func init() {
	fschedule.Schedule(
		"1 0 * * *",
		1000,
		true,
		move,
	)
	return
}

// 移动文件.
func move() {
	if "" == projectName {
		return
	}
	if nil != debugFile {
		debugFile.Close()
		debugFile = nil
	}
	if nil != infoFile {
		infoFile.Close()
		infoFile = nil
	}
	if nil != warnFile {
		warnFile.Close()
		warnFile = nil
	}
	if nil != errorFile {
		errorFile.Close()
		errorFile = nil
	}
	if nil != fatalFile {
		fatalFile.Close()
		fatalFile = nil
	}
	var time001 time.Time
	time001 = time.Now().Add(-24 * time.Hour)
	var time002 string
	time002 = fmt.Sprintf("-%02d-%02d-", time001.Month(), time001.Day())
	var a001 string
	var e error
	if F_DEBUG == F_DEBUG&logFlag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+"-DEBUG.log")
		var a002 string
		a002 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+time002+"DEBUG.log")
		os.Rename(a001, a002)
		debugFile, e = os.OpenFile(
			a001,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0664)
		if nil != e {
			panic(e)
		}
		debugWriter = io.MultiWriter(os.Stdout, debugFile)
	} else {
		debugWriter = os.Stdout
	}
	if F_INFO == F_INFO&logFlag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+"-INFO.log")
		var a002 string
		a002 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+time002+"INFO.log")
		os.Rename(a001, a002)
		infoFile, e = os.OpenFile(
			a001,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0664)
		if nil != e {
			panic(e)
		}
		infoWriter = io.MultiWriter(os.Stdout, infoFile)
	} else {
		infoWriter = os.Stdout
	}
	if F_WARN == F_WARN&logFlag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+"-WARN.log")
		var a002 string
		a002 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+time002+"WARN.log")
		os.Rename(a001, a002)
		warnFile, e = os.OpenFile(
			a001,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0664)
		if nil != e {
			panic(e)
		}
		warnWriter = io.MultiWriter(os.Stdout, warnFile)
	} else {
		warnWriter = os.Stdout
	}
	if F_ERROR == F_ERROR&logFlag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+"-ERROR.log")
		var a002 string
		a002 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+time002+"ERROR.log")
		os.Rename(a001, a002)
		errorFile, e = os.OpenFile(
			a001,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0664)
		if nil != e {
			panic(e)
		}
		errorWriter = io.MultiWriter(os.Stdout, errorFile)
	} else {
		errorWriter = os.Stdout
	}
	if F_FATAL == F_FATAL&logFlag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+"-FATAL.log")
		var a002 string
		a002 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			projectName,
			projectName+time002+"FATAL.log")
		os.Rename(a001, a002)
		fatalFile, e = os.OpenFile(
			a001,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0664)
		if nil != e {
			panic(e)
		}
		fatalWriter = io.MultiWriter(os.Stdout, fatalFile)
	} else {
		fatalWriter = os.Stdout
	}
	return
}
