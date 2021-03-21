package webframework

import (
	// TODO
	//
	"net/http"
)

// r.URL.Path
type Interceptor interface {
	Process(http.ResponseWriter, *http.Request) bool
}

// 用于注册拦截器
type InterceptorRegistry struct {
	interceptors []Interceptor
}

var InterceptorRegistryA *InterceptorRegistry

func init() {
	InterceptorRegistryA = new(InterceptorRegistry)
	InterceptorRegistryA.interceptors = make([]Interceptor, 0, 0xf)
	return
}

// 注册拦截器
func (i *InterceptorRegistry) Registry(interceptor Interceptor) {
	i.interceptors = append(i.interceptors, interceptor)
	return
}

// 处理路由, 返回true表示继续路由
func (i *InterceptorRegistry) process(w http.ResponseWriter, r *http.Request) bool {
	for _, interceptor := range i.interceptors {
		if !interceptor.Process(w, r) {
			return false
		}
	}
	return true
}
