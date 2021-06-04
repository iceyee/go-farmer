package ffile

import (
	"strings"
	//
)

// 将各个元素拼接起来合成路径.
func Path(elements ...string) string {
	return strings.Join(elements, PathSeparator)
}
