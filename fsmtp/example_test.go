package fsmtp

import (
	"github.com/iceyee/go-farmer/v4/fassert"
	"testing"
	//
)

func TestSendMail(t *testing.T) {
	var e error
	SendMail(
		QQ_MAIL_SERVER,
		"iceyee.studio@qq.com",
		"itxljrntxudubhic",
		"测试-标题",
		"测试-内容.",
		"farmer.person@qq.com",
		"709565591@qq.com",
	)
	fassert.CheckError(e)
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
		"709565591@qq.com",
	)
	fassert.CheckError(e)
	return
}
