package fsmtp

import (
	"github.com/iceyee/go-farmer/v3/fassert"
	"testing"
	//
)

func TestSendMail(t *testing.T) {
	var e error
	SendMail(
		QQ_MAIL_SERVER,
		"iceyee.studio@qq.com",
		"ddaulujsexzlbidd",
		"farmer.person@qq.com",
		"测试-标题",
		"测试-内容.",
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
		"farmer.person@qq.com",
		"测试-标题",
		"测试-内容.",
	)
	fassert.CheckError(e)
	return
}
