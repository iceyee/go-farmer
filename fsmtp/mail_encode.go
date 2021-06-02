package fsmtp

import (
	"strings"
	//
)

// 字符串转义, &<> .
func MailEncode(a string) string {
	a = strings.Replace(a, "&", "&amp;", -1)
	a = strings.Replace(a, "<", "&lt;", -1)
	a = strings.Replace(a, ">", "&gt;", -1)
	return a
}
