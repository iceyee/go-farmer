package web

import (
	"net/http"
	"strings"
	//
)

// 文件服务器, 用于处理静态文件
type FileServer struct {
	handlers []http.Handler
	prefixes []string
}

var FileServerA *FileServer

func init() {
	FileServerA = new(FileServer)
	FileServerA.handlers = make([]http.Handler, 0, 0xf)
	FileServerA.prefixes = make([]string, 0, 0xf)
	return
}

// 注册文件服务器, prefix需要以/结尾
func (f *FileServer) Registry(prefix string, directory string) {
	if !strings.HasSuffix(prefix, "/") {
		panic("必须以'/'结尾")
	}
	handler1 := http.StripPrefix(prefix, http.FileServer(http.Dir(directory)))
	f.handlers = append(f.handlers, handler1)
	f.prefixes = append(f.prefixes, prefix)
	return
}

// 处理路由, 返回true表示继续路由
func (f *FileServer) process(w http.ResponseWriter, r *http.Request) bool {
	for index, value := range f.prefixes {
		if !strings.HasPrefix(r.URL.Path, value) {
			continue
		}
		f.handlers[index].ServeHTTP(w, r)
		return false
	}
	return true
}
