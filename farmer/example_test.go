package farmer

import (
// TODO
//
)

func ExampleHttp() {
	http1 := NewHttp()
	http1.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http1.SetProxy("")
	http1.SetVerbose(true)
	http1.Request()
}
