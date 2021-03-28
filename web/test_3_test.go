package web

import (
	"encoding/base64"
	"github.com/iceyee/go-farmer/v2/farmer"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	//
)

type EncodingController struct {
}

func (c *EncodingController) GetApi() []ApiDocument {
	return []ApiDocument{
		c.a1(),
	}
}

func (*EncodingController) a1() ApiDocument {
	return ApiDocument{
		Key:          "encoding-base64",
		Url:          "/base64",
		Method:       "GET,POST",
		MapTo:        "Base64",
		ArgumentType: reflect.TypeOf(t5915{}),
		Description:  "Base64编码解码",
		Response:     "返回字符串",
	}
}

type t5915 struct {
	Source string `web:"name:source ###type:string ###require:1"`
	Method int    `web:"name:method ###type:int ###default:1 ###desc: 1表示编码, 2表示解码"`
}

func (*EncodingController) Base64(w http.ResponseWriter, r *http.Request, arg interface{}) {
	if nil == arg {
		http.Error(w, "参数解析错误", 500)
	}
	var t = arg.(t5915)

	if 1 == t.Method {
		var result1 = base64.StdEncoding.EncodeToString([]byte(t.Source))
		w.Write([]byte(result1))
	} else if 2 == t.Method {
		result1, e := base64.StdEncoding.DecodeString(t.Source)
		if nil != e {
			http.Error(w, e.Error(), 400)
		}
		w.Write(result1)
	} else {
		http.Error(w, "错误的请求", 400)
	}
	return
}

func test3(t *testing.T) {
	ControllerRegistryA.Registry(new(EncodingController))
	var server1 = httptest.NewServer(new(Server))
	defer server1.Close()

	var values = url.Values{}
	values.Add("source", "hello world.")
	values.Add("method", "1")
	var http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/base64")
	http1.SetMethod("POST")
	http1.SetBody([]byte(values.Encode()))
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())
	farmer.Assert(0 < len(http1.GetResponseBody()))
	return
}
