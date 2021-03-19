package farmer

import (
	// TODO
	//
	"fmt"
	"runtime/debug"
)

// 自定义异常, 能输出堆栈信息
type FarmerError struct {
	message string
}

// 生成异常
func NewFarmerError(message interface{}) *FarmerError {
	var message1 string
	switch message.(type) {
	case string:
		message1 = message.(string)
	case error:
		message1 = message.(error).Error()
	case fmt.Stringer:
		message1 = message.(fmt.Stringer).String()
	default:
		message1 = ""
	}
	farmerError := new(FarmerError)
	farmerError.message = message1 + "\n" + string(debug.Stack())
	return farmerError

}

func (farmerError *FarmerError) Error() string {
	return farmerError.message
}
