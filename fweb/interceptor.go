package fweb

import (
//
)

var x494 []Interceptor = make([]Interceptor, 0, 0xff)

func RegistryInterceptor(interceptor Interceptor) {
	x494 = append(x494, interceptor)
	return
}
