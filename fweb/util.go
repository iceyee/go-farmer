package fweb

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v4/ferror"
	"github.com/iceyee/go-farmer/v4/flog"
	"github.com/iceyee/go-farmer/v4/ftype"
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
func Listen(addressAndPort string) ftype.Error {
	flog.Debug("监听 " + addressAndPort)
	return ferror.New(http.ListenAndServe(addressAndPort, new(Server)))
}
