package fsmtp

import (
	"fmt"
	"github.com/iceyee/go-farmer/v5/ferror"
	"github.com/iceyee/go-farmer/v5/flog"
	"github.com/iceyee/go-farmer/v5/fstrings"
	"github.com/iceyee/go-farmer/v5/ftype"
	"net/smtp"
	"strings"
	//
)

// 发送邮件. 支持群发.
func SendMail(
	server string,
	account string,
	password string,
	subject string,
	content string,
	to ...string) ftype.Error {

	subject = strings.Replace(subject, "\n", "    ", -1)
	var sb001 *fstrings.StringBuffer
	sb001 = fstrings.NewStringBuffer()
	sb001.Append("Content-Type: text/html; charset=UTF-8; \r\n")
	sb001.Append("From: <")
	sb001.Append(account)
	sb001.Append(">\r\n")
	sb001.Append("Subject: ")
	sb001.Append(subject)
	sb001.Append("\r\n")
	sb001.Append("\r\n")
	sb001.Append(content)
	var sb002 *fstrings.StringBuffer
	sb002 = fstrings.NewStringBuffer()
	defer flog.Debug(sb002)
	sb002.Append("fsmtp.SendMail() - ")
	sb002.Append("\nserver: " + server)
	sb002.Append("\naccount: " + account)
	sb002.Append("\npassword: " + password)
	sb002.Append("\nto: " + fmt.Sprintf("%v", to))
	sb002.Append("\nsubject: " + subject)
	sb002.Append("\ncontent: " + content)
	sb002.Append("\n\n>>>>>>>>>>\n")
	sb002.Append(sb001)
	var auth001 smtp.Auth
	if strings.Contains(server, ":") {
		auth001 = smtp.PlainAuth(
			"",
			account,
			password,
			strings.Split(server, ":")[0])
	} else {
		auth001 = smtp.PlainAuth(
			"",
			account,
			password,
			server)
	}
	var e error
	e = smtp.SendMail(
		server,
		auth001,
		account,
		to,
		[]byte(sb001.String()))
	if nil != e {
		e = ferror.New(e)
	}
	return e
}
