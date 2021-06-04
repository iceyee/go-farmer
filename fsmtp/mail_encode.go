package fsmtp

import (
	"strings"
	//
)

// 字符串转义, &<>\n .
// 不要对同一字符串转义两次, 会出现非预期的情况.
func MailEncode(a string) string {
	a = strings.Replace(a, "&", "&amp;", -1)
	a = strings.Replace(a, "<", "&lt;", -1)
	a = strings.Replace(a, ">", "&gt;", -1)
	a = strings.Replace(a, "\n", "<br>", -1)
	return a
}
