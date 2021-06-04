package fconsole

import (
	"strings"
	"unicode/utf8"
	//
)

// 创建菜单.
func CreateMenu(items ...string) string {
	if 0 == len(items) {
		return ""
	}
	var maxLength int
	for _, x := range items {
		if maxLength < getLength(x) {
			maxLength = getLength(x)
		}
	}
	var a001 string
	a001 += strings.Repeat("#", 5+maxLength+5)
	for _, x := range items {
		a001 += "\n"
		a001 += "#    "
		a001 += x
		a001 += strings.Repeat(" ", maxLength-getLength(x))
		a001 += "    #"
	}
	a001 += "\n"
	a001 += strings.Repeat("#", 5+maxLength+5)
	return a001
}

// 中文是一个字占三个字节, 显示宽度占两个字母的宽度.
func getLength(content string) int {
	var a001 int
	var a002 int
	a001 = len(content)
	a002 = utf8.RuneCountInString(content)
	return (a001 + a002) / 2
}
