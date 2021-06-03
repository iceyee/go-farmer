package fweb

import (
	"github.com/iceyee/go-farmer/v4/flog"
	"reflect"
	//
)

var x494 []Interceptor = make([]Interceptor, 0, 0xff)

// 注册拦截器.
func RegistryInterceptor(interceptor Interceptor) {
	x494 = append(x494, interceptor)
	flog.Debug("Interceptor, " + reflect.TypeOf(interceptor).String())
	return
}
