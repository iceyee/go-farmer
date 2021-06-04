package fschedule

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	//
)

// 定时任务, 循环执行.
// 例如: 0 5 * * * , 这表示每天5点执行.
// 第五个数, 0表示星期天, 其它1-6表示对应周几.
// @desc - 对时间的描述, 参考cron.
// @delay - 初始延时, 单位毫秒.
// @repeated - 是否重复任务, false表示只执行一次.
// @f - 任务.
func Schedule(
	desc string,
	delay int64,
	repeated bool,
	f func()) {

	go func() {
		if nil == f {
			panic("没有指定任务.")
		}
		if delay < 0 {
			panic("延时不能为负数.")
		}
		for strings.Contains(desc, "  ") {
			desc = strings.Replace(desc, "  ", " ", -1)
		}
		var a001 []string
		a001 = strings.Split(desc, " ")
		if 5 != len(a001) {
			panic("Schedule语法错误.")
		}
		var b001 []int64
		var b002 []int64
		var b003 []int64
		var b004 []int64
		var b005 []int64
		f001 := func(min int64, max int64, text string) []int64 {
			text = strings.Replace(text, "*", fmt.Sprintf("%d-%d", min, max), -1)
			var a002 []int64
			a002 = make([]int64, 0, 0xfff)
			var a003 []string
			a003 = strings.Split(text, ",")
			for _, x := range a003 {
				if "*" == x {
					for y := min; y <= max; y++ {
						a002 = append(a002, y)
					}
				} else if ok, _ := regexp.MatchString(`^\d{1,2}$`, x); ok {
					var a004 int64
					a004, _ = strconv.ParseInt(x, 10, 64)
					a002 = append(a002, a004)
				} else if ok, _ := regexp.MatchString(`^\d{1,2}-\d{1,2}$`, x); ok {
					var reg *regexp.Regexp
					var e error
					reg, e = regexp.Compile(`^(\d{1,2})-(\d{1,2})$`)
					if nil != e {
						panic(e)
					}
					var a005 []string
					a005 = reg.FindStringSubmatch(x)
					if 0 == len(a005) {
						panic("Schedule语法错误.")
					}
					var a006 int64
					var a007 int64
					a006, _ = strconv.ParseInt(a005[1], 10, 64)
					a007, _ = strconv.ParseInt(a005[2], 10, 64)
					for y := a006; y <= a007; y++ {
						a002 = append(a002, y)
					}
				} else if ok, _ := regexp.MatchString(`^\d{1,2}-\d{1,2}/\d{1,2}$`, x); ok {
					var reg *regexp.Regexp
					var e error
					reg, e = regexp.Compile(`^(\d{1,2})-(\d{1,2})/(\d{1,2})$`)
					if nil != e {
						panic(e)
					}
					var a005 []string
					a005 = reg.FindStringSubmatch(x)
					if 0 == len(a005) {
						panic("Schedule语法错误.")
					}
					var a006 int64
					var a007 int64
					var a008 int64
					a006, _ = strconv.ParseInt(a005[1], 10, 64)
					a007, _ = strconv.ParseInt(a005[2], 10, 64)
					a008, _ = strconv.ParseInt(a005[3], 10, 64)
					for y := a006; y <= a007; {
						a002 = append(a002, y)
						y += a008
					}
				} else {
					panic("Schedule语法错误.")
				}
			}
			return a002
		}
		b001 = f001(0, 59, a001[0])
		b002 = f001(0, 23, a001[1])
		b003 = f001(1, 31, a001[2])
		b004 = f001(1, 12, a001[3])
		b005 = f001(0, 6, a001[4])
		// fmt.Printf("%v\n", b001)
		// fmt.Printf("%v\n", b002)
		// fmt.Printf("%v\n", b003)
		// fmt.Printf("%v\n", b004)
		// fmt.Printf("%v\n", b005)
		f002 := func(a []int64, b int64) bool {
			for _, x := range a {
				if x == b {
					return true
				}
			}
			return false
		}
		time.Sleep(time.Duration(delay) * time.Millisecond)
		for true {
			var time001 time.Time
			time001 = time.Now()
			if f002(b001, int64(time001.Minute())) &&
				f002(b002, int64(time001.Hour())) &&
				f002(b003, int64(time001.Day())) &&
				f002(b004, int64(time001.Month())) &&
				f002(b005, int64(time001.Weekday())) {
				// 满足条件, 开始执行任务
				go f()
			}
			// println(time001.Minute())
			// println(time001.Hour())
			// println(time001.Day())
			// println(time001.Month())
			// println(time001.Weekday())
			// 下一分钟
			time.Sleep(
				time.Duration(
					61-int64(time001.Second())) * time.Second)
		}
	}()
	return
}
