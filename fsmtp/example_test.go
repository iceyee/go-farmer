package fsmtp

import (
	"github.com/iceyee/go-farmer/v5/fassert"
	"testing"
	//
)

func TestMailEncode(t *testing.T) {
	fassert.Assert("1&amp;2&lt;br&gt;" == MailEncode("1&2<br>"))
	return
}

func TestSendMail(t *testing.T) {
	var e error
	SendMail(
		QQ_MAIL_SERVER,
		"iceyee.studio@qq.com",
		"itxljrntxudubhic",
		"测试-标题",
		"测试-内容.",
		"farmer.person@qq.com",
		"709565591@qq.com")
	fassert.CheckError(e, "发送邮件.")
	return
}

func ExampleSendMail() {
	var e error
	SendMail(
		QQ_MAIL_SERVER,
		"iceyee.studio@qq.com",
		"ddaulujsexzlbidd",
		"测试-标题",
		"测试-内容.",
		"farmer.person@qq.com",
		"709565591@qq.com")
	fassert.CheckError(e, "发送邮件.")
	return
}
