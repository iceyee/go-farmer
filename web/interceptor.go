package webframework

import (
	// TODO
	//
	"net/http"
	"strings"
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

	domainInterceptorA = new(domainInterceptor)
	domainInterceptorA.domains = make([]string, 0, 0xf)
	domainInterceptorA.domains = append(domainInterceptorA.domains, "localhost")
    domainInterceptorA.domains = append(domainInterceptorA.domains, "127.0.0.1")
	InterceptorRegistryA.Registry(domainInterceptorA)
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

type domainInterceptor struct {
	domains []string
}

var domainInterceptorA *domainInterceptor

func (i *domainInterceptor) Process(w http.ResponseWriter, r *http.Request) bool {
	r.Host = strings.ToLower(r.Host)
	for _, value := range i.domains {
		if strings.Contains(r.Host, value) {
			return true
		}
	}
	http.Error(w, "非法请求", 403)
	return false
}

// 授权域名, 允许通过指定域名访问, 默认只允许通过localhost访问.
func AuthorizeDomain(domain string) {
	if strings.Contains(strings.Join(domainInterceptorA.domains, ","), domain) {
		return
	}
	domainInterceptorA.domains = append(domainInterceptorA.domains, domain)
	return
}
