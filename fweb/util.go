package fweb

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/flog"
	"github.com/iceyee/go-farmer/v3/ftype"
	"net/http"
	//
)

// 向客户端输出json.
func WriteJson(w http.ResponseWriter, data interface{}) ftype.Error {
	content, e := json.Marshal(data)
	if nil != e {
		return ferror.New(e)
	}
	_, e = w.Write(content)
	if nil != e {
		return ferror.New(e)
	}
	w.Header().Set("Content-Type", "application/json")
	return nil
}

// 启动服务器
func Listen(addressAndPort string) error {
	flog.Debug("监听 " + addressAndPort)
	return http.ListenAndServe(addressAndPort, new(Server))
}
