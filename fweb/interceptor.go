package fweb

import (
	"github.com/iceyee/go-farmer/v5/flog"
	"reflect"
	//
)

var x494 []Interceptor = make([]Interceptor, 0, 0xff)

// 注册拦截器.
func RegistryInterceptor(interceptor Interceptor) {
	x494 = append(x494, interceptor)
	flog.Debug("RegistryInterceptor, " + reflect.TypeOf(interceptor).String())
	return
}
