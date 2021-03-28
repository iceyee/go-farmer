package web

import (
	"encoding/json"
	"github.com/iceyee/go-farmer/v2/farmer"
	"net/http"
	//
)

// 向客户端输出Json数据, data是一个指针(*struct)
func WriteJson(w http.ResponseWriter, data interface{}) error {
	// bytes1 - data经由json编码得到([]byte)
	bytes1, e := json.Marshal(data)
	if nil != e {
		return farmer.NewFarmerError(e)
	}

	w.Header().Set("Content-Type", "application/json")
	_, e = w.Write(bytes1)
	if nil != e {
		return farmer.NewFarmerError(e)
	}

	return nil
}
