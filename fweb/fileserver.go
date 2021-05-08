package fweb

import (
	"net/http"
	"strings"
	//
)

var x766 []t233 = make([]t233, 0, 0xff)

// 注册文件服务器, prefix需要以'/'结尾.
func RegistryFileServer(prefix string, directory string) {
	if !strings.HasSuffix(prefix, "/") {
		panic("必须以'/'结尾.")
	}
	a001 := t233{
		Handler: http.StripPrefix(prefix, http.FileServer(http.Dir(directory))),
		Prefix:  prefix,
	}
	x766 = append(x766, a001)
	return
}
