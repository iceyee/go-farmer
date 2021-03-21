package webframework

import (
	// TODO
	//
	"net/http"
)

// 启动服务器
func Listen(addressAndPort string) error {
	return http.ListenAndServe(addressAndPort, new(server))
}
