package fhttp

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v5/fassert"
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
	var http *Http
	http = New()
	http.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http.SetProxy("socks5://farmer:74591870@hk.farmer.ink:10002")
	fassert.Assert(http.Request(), "访问成功.")
	fassert.Assert(200 == http.GetStatusCode(), "返回200.")
	fassert.Assert(nil != http.GetResponseHeader(), "有响应头.")
	var a t_1244
	fassert.CheckError(json.Unmarshal(http.GetResponseBody(), &a), "json解码.")
	return
}

func ExampleHttp() {
	var http *Http
	http = New()
	http.SetUrl("http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json")
	http.SetProxy("socks5://farmer:74591870@hk.farmer.ink:10002")
	fassert.Assert(http.Request(), "访问成功.")
	fassert.Assert(200 == http.GetStatusCode(), "返回200.")
	fassert.Assert(nil != http.GetResponseHeader(), "有响应头.")
	var a t_1244
	fassert.CheckError(json.Unmarshal(http.GetResponseBody(), &a), "json解码.")
	return
}
