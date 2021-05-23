package fconsole

import (
	"strings"
	//
)

// 创建菜单.
func CreateMenu(items ...string) string {
	if 0 == len(items) {
		return ""
	}
	var maxLength int
	for _, x := range items {
		if maxLength < len(x) {
			maxLength = len(x)
		}
	}
	var a001 string
	a001 += strings.Repeat("#", 5+maxLength+5)
	for _, x := range items {
		a001 += "\n"
		a001 += "#    "
		a001 += x
		a001 += strings.Repeat(" ", maxLength-len(x))
		a001 += "    #"
	}
	a001 += "\n"
	a001 += strings.Repeat("#", 5+maxLength+5)
	return a001
}
