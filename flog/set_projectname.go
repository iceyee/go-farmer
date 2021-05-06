package flog

import (
	"github.com/iceyee/go-farmer/v3/ffile"
	"io"
	"os"
	//
)

func init() {
	ffile.InstallDirectory()
	ffile.Mkdir(ffile.Path(ffile.HomeDirectory, "farmer", "share", "farmer-log"))
	return
}

// 允许日志记录到文件中.
// @name - 项目名.
// @flag - 标志, 表示哪个日志级别需要记录.
func SetProjectName(name string, flag LogFlag) {
	if "" == name {
		panic("ProjectName不能为空.")
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
	ffile.Mkdir(ffile.Path(ffile.HomeDirectory, "farmer", "share", "farmer-log", name))
	var a001 string
	var e error
	if F_DEBUG == F_DEBUG&flag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			name,
			name+"-DEBUG.log")
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
	if F_INFO == F_INFO&flag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			name,
			name+"-INFO.log")
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
	if F_WARN == F_WARN&flag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			name,
			name+"-WARN.log")
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
	if F_ERROR == F_ERROR&flag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			name,
			name+"-ERROR.log")
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
	if F_FATAL == F_FATAL&flag {
		a001 = ffile.Path(
			ffile.HomeDirectory,
			"farmer",
			"share",
			"farmer-log",
			name,
			name+"-FATAL.log")
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
	projectName = name
	logFlag = flag
	return
}
