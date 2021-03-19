package webframework

import (
	// TODO
	//
	"github.com/iceyee/go-farmer/farmer"
	"net/http"
	"testing"
)

// 这个拦截器, 添加Cookie: a=b; c=d
type Interceptor1 struct{}

func (i *Interceptor1) Process(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Add("Set-Cookie", "a=b; path=/")
	w.Header().Add("Set-Cookie", "c=d; path=/")
	return true
}

// 这个拦截器, 禁止访问/tmp/vim-paste1
type Interceptor2 struct{}

func (i *Interceptor2) Process(w http.ResponseWriter, r *http.Request) bool {
	if "/tmp/vim-paste1" == r.URL.Path ||
		"/tmp/vim-paste1/" == r.URL.Path {
		http.Error(w, "禁止访问这个文件", 403)
		return false
	}
	return true
}

type Controller1 struct{}

// /hello, 返回hello world.
func (c *Controller1) Hello(w http.ResponseWriter, r *http.Request) {
	_, e := w.Write([]byte("hello world"))
	if nil != e {
		panic(e)
	}
	return
}

type T1 struct {
	A int     `json:"a"`
	B float32 `json:"b"`
	C string  `json:"c"`
}

// /json, 测试json格式的数据
func (c *Controller1) Json(w http.ResponseWriter, r *http.Request) {
	var t1 T1
	t1.A = 1
	t1.B = 0.2
	t1.C = "hello world"
	e := WriteJson(w, &t1)
	farmer.Assert(nil == e)
	return
}

type T0 struct {
	A int     `name:"a" require:"true" enum:"1,2,4,8"`
	B float32 `name:"B" require:"true" min:"0.1" max:"0.9"`
	C string  `name:"c" default:"hello world" enum:"fg,we"`
	D int16   `name:"d" min:"1" max:"99"`
	E string
}

// 测试Validate表单验证的功能
func (c *Controller1) Validate(w http.ResponseWriter, r *http.Request) {
	var t0 T0
	valid, e := Validate(w, r, &t0)
	if nil != e {
		println(e.Error())
	}
	if valid {
		println("表单有效")
	} else {
		println("表单无效")
	}
	println("T0:")
	println(t0.A)
	println(t0.B)
	println(t0.C)
	println(t0.D)
	println(t0.E)
	println()
	return
}

func (c *Controller1) Validate_api() ApiDocument {
	var a ApiDocument
	a.Description = "测试Validate表单验证的功能"
	a.Key = "1"
	a.Method = "GET, POST"
	a.Parameters = "a(必须): [1,2,4,8] 整数, B(必须): 0.1~0.9 小数, c: ['fg','we'], d: 1~99 整数"
	a.Remarks = "无"
	a.Response = ""
	a.Url = "/validate"
	return a
}

func TestServer(t *testing.T) {
	AcceptMethod("DELETE")
	InterceptorRegistryA.Registry(new(Interceptor1))
	InterceptorRegistryA.Registry(new(Interceptor2))
	ControllerRegistryA.Registry(new(Controller1))
	FileServerA.Registry("/tmp/", "/tmp/")
    Listen(":8888")
	return
}
