package webframework

import (
	// TODO
	//
	"net/http"
)

// 处理请求方法的拦截器, 默认只处理GET和POST
type MethodInterceptor struct {
	Method map[string]bool
}

var MethodInterceptorA *MethodInterceptor

func (m *MethodInterceptor) Process(w http.ResponseWriter, r *http.Request) bool {
	if !m.Method[r.Method] {
		http.Error(w, "禁止的请求方法", 405)
		return false
	}
	return true
}

// 允许指定的请求方法
func AcceptMethod(method string) {
	MethodInterceptorA.Method[method] = true
	return
}

// 禁止指定的请求方法
func ForbidMethod(method string) {
	MethodInterceptorA.Method[method] = false
	return
}
