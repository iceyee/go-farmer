package farmer

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	//
)

// project - 项目名(string)
// running - 控制监控的行为, 为true的时候表示通过(bool)
var fileERROR *os.File
var fileINFO *os.File
var fileWARN *os.File
var loggerERROR *log.Logger
var loggerINFO *log.Logger
var loggerWARN *log.Logger
var project string
var running bool

func init() {
	loggerINFO = log.New(io.MultiWriter(os.Stdout), "\nINFO   ", log.LstdFlags)
	loggerWARN = log.New(io.MultiWriter(os.Stdout), "\nWARN   ", log.LstdFlags)
	loggerERROR = log.New(io.MultiWriter(os.Stderr), "\nERROR  ", log.LstdFlags)
	go deleteOldFile()
	go move()
	go synchronize()
	return
}

// 开启日志系统, 将Log记录到指定目录
func OpenLog(projectName string) {
	// fileInfo1 - /opt/farmer-log/ 目录信息(os.FileInfo)
	// name1 - /opt/farmer-log/xxx/ 目标目录, 保存日志的位置(string)
	// fileInfo2 - /opt/farmer-log/xxx/ 目录信息(os.FileInfo)
	fileInfo1, e := os.Stat("/opt/farmer-log/")
	if nil != e || !fileInfo1.IsDir() {
		panic("缺少'/opt/farmer-log/'目录")
	}
	project = projectName
	name1 := "/opt/farmer-log/" + project
	fileInfo2, e := os.Stat(name1)
	if nil != e {
		e = os.Mkdir(name1, os.ModeDir|0755)
		if nil != e {
			panic(e)
		}
	} else if !fileInfo2.IsDir() {
		e = os.Remove(name1)
		if nil != e {
			panic(e)
		}
		e = os.Mkdir(name1, os.ModeDir|0644)
		if nil != e {
			panic(e)
		}
	}

	// warn1 - /opt/farmer-log/xxx/xxx-WARN.log (string)
	// error1 - /opt/farmer-log/xxx/xxx-ERROR.log (string)
	// fileWARN, fileERROR - (os.File)
	// loggerINFO, loggerWARN, loggerERROR - (log.Logger)
	loggerINFO = log.New(io.MultiWriter(os.Stdout), "\nINFO   ", log.LstdFlags)
	warn1 := fmt.Sprintf("/opt/farmer-log/%s/%s-WARN.log", project, project)
	fileWARN, e = os.OpenFile(warn1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if nil != e {
		panic(e)
	}
	loggerWARN = log.New(io.MultiWriter(fileWARN, os.Stdout), "\nWARN   ", log.LstdFlags)
	error1 := fmt.Sprintf("/opt/farmer-log/%s/%s-ERROR.log", project, project)
	fileERROR, e = os.OpenFile(error1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if nil != e {
		panic(e)
	}
	loggerERROR = log.New(io.MultiWriter(fileERROR, os.Stderr), "\nERROR  ", log.LstdFlags)

	fmt.Printf("OPENLOG('%s')!!!\n", project)

	running = true
	return
}

// 关闭日志系统
func CloseLog() {
	if nil != fileINFO {
		fileINFO.Close()
	}

	if nil != fileWARN {
		fileWARN.Close()
	}

	if nil != fileERROR {
		fileERROR.Close()
	}

	running = false
	return
}

func Info(message interface{}) {
	switch message.(type) {
	case string:
		loggerINFO.Println(message.(string))
	case error:
		loggerINFO.Println(message.(error).Error())
	case fmt.Stringer:
		loggerINFO.Println(message.(fmt.Stringer).String())
	default:
		loggerINFO.Println("[Unkown]")
	}
	return
}

func Warn(message interface{}) {
	switch message.(type) {
	case string:
		loggerWARN.Println(message.(string))
	case error:
		loggerWARN.Println(message.(error).Error())
	case fmt.Stringer:
		loggerWARN.Println(message.(fmt.Stringer).String())
	default:
		loggerWARN.Println("[Unkown]")
	}
	return
}

func Error(message interface{}) {
	switch message.(type) {
	case string:
		loggerERROR.Println(message.(string))
	case error:
		loggerERROR.Println(message.(error).Error())
	case fmt.Stringer:
		loggerERROR.Println(message.(fmt.Stringer).String())
	default:
		loggerERROR.Println("[Unkown]")
	}
	return
}

// 删除7天前的日志
func deleteOldFile() {
	time.Sleep(1 * time.Second)
	for true {
		if !running {
			time.Sleep(1 * time.Minute)
			continue
		}
		func() {
			// directory1 - /opt/farmer-log/项目名(string)
			// state1 - /opt/farmer-log/项目名(os.FileInfo)
			// now1 - 现在的时间戳, 单位:纳秒(int64)
			// files1 - 一组文件信息([]os.FileInfo)
			// file - /opt/farmer-log/项目名/xxx (os.FileInfo)
			if "" == project {
				return
			}
			directory1, e := os.Open("/opt/farmer-log/" + project)
			if nil != e {
				panic(e)
			}
			state1, e := directory1.Stat()
			if nil != e {
				panic(e)
			} else if !state1.IsDir() {
				panic("Expected direcotry")
			}
			now1 := time.Now().UnixNano()
			files1, e := directory1.Readdir(0)
			for _, file := range files1 {
				if file.IsDir() {
					continue
				} else if 8*24*time.Hour.Nanoseconds() < now1-file.ModTime().UnixNano() {
					os.Remove(file.Name())
					os.Remove(fmt.Sprintf("/opt/farmer-log/%s/%s", project, file.Name()))
				}
			}
		}()
		time.Sleep(1 * time.Minute)
	}
	return
}

// 0:01 日志文件转移
func move() {
	time.Sleep(1 * time.Second)
	for true {
		if !running {
			time.Sleep(1 * time.Minute)
			continue
		}
		func() {
			// now1 - 现在的时间(time.Time)
			// date1 - 日期, 格式: 月-日, M-dd (string)
			// warn1 - /opt/farmer-log/项目名/项目名-WARN.log (string)
			// warn2 - /opt/farmer-log/项目名/项目名-WARN-日期.log (string)
			// error1 - /opt/farmer-log/项目名/项目名-ERROR.log (string)
			// error2 - /opt/farmer-log/项目名/项目名-ERROR-日期.log (string)
			now1 := time.Now()
			if "" == project {
				return
			} else if 0 != now1.Hour() || 1 != now1.Minute() {
				return
			}

			loggerINFO.SetOutput(os.Stdout)
			loggerWARN.SetOutput(os.Stdout)
			loggerERROR.SetOutput(os.Stdout)

			now1 = now1.Add(-24 * time.Hour)
			date1 := fmt.Sprintf("%d-%2d", now1.Month(), now1.Day())

			fileWARN.Close()
			fileERROR.Close()

			warn1 := fmt.Sprintf("/opt/farmer-log/%s/%s-WARN.log", project, project)
			warn2 := fmt.Sprintf("/opt/farmer-log/%s/%s-WARN-%s.log", project, project, date1)
			error1 := fmt.Sprintf("/opt/farmer-log/%s/%s-ERROR.log", project, project)
			error2 := fmt.Sprintf("/opt/farmer-log/%s/%s-ERROR-%s.log", project, project, date1)

			os.Rename(warn1, warn2)
			os.Rename(error1, error2)

			var e error
			loggerINFO = log.New(io.MultiWriter(os.Stdout), "\nINFO   ", log.LstdFlags)
			fileWARN, e = os.OpenFile(warn1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if nil != e {
				panic(e)
			}
			loggerWARN = log.New(io.MultiWriter(fileWARN, os.Stdout), "\nWARN   ", log.LstdFlags)
			fileERROR, e = os.OpenFile(error1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if nil != e {
				panic(e)
			}
			loggerERROR = log.New(io.MultiWriter(fileERROR, os.Stderr), "\nERROR  ", log.LstdFlags)

			time.Sleep(60 * time.Second)
		}()
		time.Sleep(59 * time.Second)
	}
	return
}

// 把日志写到磁盘
func synchronize() {
	time.Sleep(1 * time.Second)
	for true {
		if !running {
			time.Sleep(1 * time.Minute)
			continue
		}
		func() {
			if "" == project {
				return
			}
			fileWARN.Sync()
			fileERROR.Sync()
		}()
		time.Sleep(1 * time.Minute)
	}
	return
}
