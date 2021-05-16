package fhttp

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v4/fassert"
	"testing"
	//
)

type t_1244 struct {
	Success string `json:"success"`
	Result  struct {
		Ip        string `json:"ip"`
		Proxy     string `json:"proxy"`
		Attribute string `json:"att"`
		Operators string `json:"operators"`
	} `json:"result"`
}

func TestHttp(t *testing.T) {
	var http001 *Http
	http001 = New()
	http001.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http001.SetProxy("socks5://farmer:74591870@hk.farmer.ink:10002")
	fassert.Assert(http001.Request(), "访问成功")
	fassert.Assert(200 == http001.GetStatusCode(), "返回200")
	fassert.Assert(nil != http001.GetResponseHeader(), "有响应头")
	var a t_1244
	fassert.CheckError(json.Unmarshal(http001.GetResponseBody(), &a))
	return
}

func ExampleHttp() {
	var http001 *Http
	http001 = New()
	http001.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http001.SetProxy("socks5://farmer:74591870@hk.farmer.ink:10002")
	fassert.Assert(http001.Request(), "访问成功")
	fassert.Assert(200 == http001.GetStatusCode(), "返回200")
	fassert.Assert(nil != http001.GetResponseHeader(), "有响应头")
	var a t_1244
	fassert.CheckError(json.Unmarshal(http001.GetResponseBody(), &a))
	return
}
