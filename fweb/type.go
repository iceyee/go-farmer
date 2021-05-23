package fweb

import (
	"net/http"
	"reflect"
	//
)

// Url参数的定义.
type t377 struct {
	Default     string
	Description string
	Max         *float64
	Min         *float64
	Name        string
	Not         string
	Regexp      string
	Required    bool
	Type        string
}

// 单个web接口的定义.
type t184 struct {
	Description string
	Executor    reflect.Value
	MapTo       reflect.Value
	Method      string
	Parameters  []t377
	Remarks     string
	Response    string
	SortKey     string
	Url         string
}

// 文件服务器.
type t233 struct {
	Handler http.Handler
	Prefix  string
}

type Controller interface {
}

// 拦截器.
type Interceptor interface {
	// 过滤规则, 返回true表示执行.
	Filter(path string) bool
	// 处理请求.
	Process(session *Session, w http.ResponseWriter, r *http.Request) bool
}

// 可以作为Restful Api的返回结果.
type T404 struct {
	Data    interface{}
	Message string
	Result  bool
}
