package webframework

import (
	// TODO
	//
	"encoding/json"
	"github.com/iceyee/go-farmer/farmer"
	"net/http"
)

// 给客户端返回Json数据
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
