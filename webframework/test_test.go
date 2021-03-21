package webframework

import (
	// TODO
	//
	// "github.com/iceyee/go-farmer/farmer"
	"net/http"
	"reflect"
	"testing"
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
	A string `web:" name:A ### require: ### description:描述1 ### type:string ### not:hello "`
	B int64  `web:" name:b ### description:这个要求是十六进制数 ### type:hex ### max:0xff ### min:0x2 ### not:50 ### default:0xf"`
	C bool   `web:" name:c ### type:bool ### default:false "`
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
	print("A=")
	println(a1.A)
	print("b=")
	println(a1.B)
	print("c=")
	println(a1.C)
	println()
	println()
	w.Write([]byte("ok"))
	return
}

func Test(t *testing.T) {
    AuthorizeDomain("farmer.ink")
	InterceptorRegistryA.Registry(new(farmerInterceptor))
	ControllerRegistryA.Registry(new(farmerController))
	FileServerA.Registry("/tmp/", "/tmp/")
    // Listen(":8888")
	return
}
