package ferror

import (
	"github.com/iceyee/go-farmer/v5/ftype"
	"runtime/debug"
	//
)

type FarmerError struct {
	message string
}

// @message - 支持类型包括, string, error, ftype.Stringer, []byte.
func New(message interface{}) *FarmerError {
	var e *FarmerError
	e = new(FarmerError)
	switch message.(type) {

	case string:
		e.message = message.(string)
	case error:
		e.message = message.(error).Error()
	case ftype.Stringer:
		e.message = message.(ftype.Stringer).String()
	case []byte:
		e.message = string(message.([]byte))
	default:
		e.message = "[UNKWON MESSAGE]"
	}
	e.message += "\n"
	e.message += string(debug.Stack())
	return e
}

func (f *FarmerError) Error() string {
	return f.message
}
