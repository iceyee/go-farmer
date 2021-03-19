package farmer

import (
	// TODO
	//
	"testing"
)

func testHttp(t *testing.T) {
	http1 := NewHttp()
	http1.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http1.SetProxy("socks5://farmer:74591870@hk.farmer.ink:10002")
	http1.SetVerbose(true)
	Assert(http1.Request())
	println(http1.GetStatusCode())
	println(http1.GetResponseHeader())
	println(string(http1.GetResponseBody()))
	return
}
