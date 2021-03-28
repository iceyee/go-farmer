package web

import (
	"github.com/iceyee/go-farmer/v1/farmer"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	//
)

// 这个拦截器, 把HEAD请求方法改成GET
type farmerInterceptor struct {
}

func (i *farmerInterceptor) Process(w http.ResponseWriter, r *http.Request) bool {
	if "HEAD" == r.Method {
		r.Method = "GET"
	}
	return true
}

type farmerController struct {
}

func (c *farmerController) GetApi() []ApiDocument {
	return []ApiDocument{
		c.a1(),
	}
}

type t4715 struct {
	A string `web:"name:A ###type:string ###require: ###desc:描述1 ###not:hello "`
	B int64  `web:"name:b ###type:hex ###default:0xf ###desc:这个要求是十六进制数 ###max:0xff ###min:0x2 ###not:50"`
	C bool   `web:"name:c ###type:bool ###default:false "`
}

func (c *farmerController) a1() ApiDocument {
	return ApiDocument{
		ArgumentType: reflect.TypeOf(t4715{}),
		Description:  "这是第一个测试接口",
		Key:          "test1",
		MapTo:        "A",
		Method:       "GET,POST",
		Remarks:      "测试1",
		Response:     "如果参数验证没问题, 就返回ok, 否则返回错误的参数, 以及状态码400",
		Url:          "/0/test1",
	}
}

func (c *farmerController) A(w http.ResponseWriter, r *http.Request, arg interface{}) {
	a1 := arg.(t4715)
	var stringBuilder1 = farmer.NewStringBuilder()
	stringBuilder1.Append("A=")
	stringBuilder1.Append(a1.A)
	stringBuilder1.Append(", B=")
	stringBuilder1.Append(a1.B)
	stringBuilder1.Append(", C=")
	stringBuilder1.Append(a1.C)
	stringBuilder1.Append("\n")
	// println(stringBuilder1.String())
	w.Write([]byte(stringBuilder1.String()))
	return
}

func test1(t *testing.T) {
	InterceptorRegistryA.Registry(new(farmerInterceptor))
	ControllerRegistryA.Registry(new(farmerController))
	FileServerA.Registry("/tmp/", "/tmp/")
	// Listen(":8888")

	// 开始测试
	var server1 = httptest.NewServer(new(Server))
	defer server1.Close()

	// 测试域名访问的功能, 不允许test.farmer.ink
	var http1 = farmer.NewHttp()
	http1.SetUrl(strings.Replace(server1.URL, "127.0.0.1", "test.farmer.ink", -1))
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(403 == http1.GetStatusCode() || 500 == http1.GetStatusCode())
	// 添加域名, 访问成功
	AuthorizeDomain("farmer.ink")
	http1 = farmer.NewHttp()
	http1.SetUrl(strings.Replace(server1.URL, "127.0.0.1", "test.farmer.ink", -1))
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())

	// 测试文件服务, 写入/tmp/hello, 访问通过
	farmer.Assert(nil == ioutil.WriteFile("/tmp/hello", []byte("hello world"), 0644))
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/tmp/hello")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())
	// 访问/tmp/hello/hello, 不存在返回Not Found
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/tmp/hello/hello")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(404 == http1.GetStatusCode())

	// 测试拦截器, 访问/0/api, PUT不允许
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/api")
	http1.SetMethod("PUT")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(405 == http1.GetStatusCode())
	// HEAD通过, 被拦截器改GET, 通过
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/api")
	http1.SetMethod("HEAD")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())
	// ioutil.WriteFile("/tmp/1html.html", http1.GetResponseBody(), 0644)
	// farmer.Assert(nil == exec.Command("/usr/bin/firefox", "/tmp/1html.html").Run())

	// 测试控制器, 测试Url参数
	println("不带A, 错误")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?b=1f&c=true")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(400 == http1.GetStatusCode())

	println("b不是十六进制, 错误")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer&b=gg&c=true")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(400 == http1.GetStatusCode())

	println("c不是bool类型, 错误")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer&b=0x1f&c=1")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(400 == http1.GetStatusCode())

	println("c不是bool类型, 错误")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer&b=0x1f&c=1")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(400 == http1.GetStatusCode())

	println("正常访问")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer&b=0x1f&c=true")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())

	println("正常访问, b和c默认")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())

	println("b太小, 错误")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer&b=0")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(400 == http1.GetStatusCode())

	println("b太大, 错误")
	http1 = farmer.NewHttp()
	http1.SetUrl(server1.URL + "/0/test1?A=farmer&b=0xffff")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(400 == http1.GetStatusCode())

	return
}
