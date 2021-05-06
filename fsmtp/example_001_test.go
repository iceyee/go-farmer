package fsmtp

import (
	"github.com/iceyee/go-farmer/v3/fassert"
	//
)

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
