package farmer

import (
	"encoding/json"
	"testing"
	//
)

type t_1244 struct {
	Success string   `json:"success"`
	Result  t_1_1244 `json:"result"`
}

type t_1_1244 struct {
	Ip        string `json:"ip"`
	Proxy     string `json:"proxy"`
	Attribute string `json:"att"`
	Operators string `json:"operators"`
}

func testHttp(t *testing.T) {
	http1 := NewHttp()
	http1.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http1.SetProxy("socks5://farmer:74591870@hk.farmer.ink:10002")
	http1.SetVerbose(true)
	Assert(http1.Request())
	Assert(200 == http1.GetStatusCode())
	Assert(nil != http1.GetResponseHeader())
	var a t_1244
	Assert(nil == json.Unmarshal(http1.GetResponseBody(), &a))
	return
}
