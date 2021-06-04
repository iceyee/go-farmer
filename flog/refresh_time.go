package flog

import (
	"fmt"
	"time"
	//
)

func init() {
	func() {
		var time001 time.Time
		time001 = time.Now()
		var a001 *string
		a001 = new(string)
		*a001 =
			fmt.Sprintf(
				"%02d/%02d/%02d %02d:%02d:%02d",
				time001.Year(),
				time001.Month(),
				time001.Day(),
				time001.Hour(),
				time001.Minute(),
				time001.Second(),
			)
		time355 = a001
	}()
	go refreshTime()
	return
}

// 刷新显示的时间.
func refreshTime() {
	for true {
		time.Sleep(33 * time.Millisecond)
		go func() {
			var time001 time.Time
			time001 = time.Now()
			var a001 *string
			a001 = new(string)
			*a001 =
				fmt.Sprintf(
					"%02d/%02d/%02d %02d:%02d:%02d",
					time001.Year(),
					time001.Month(),
					time001.Day(),
					time001.Hour(),
					time001.Minute(),
					time001.Second())
			time355 = a001
		}()
	}
	return
}
