package farmer

import (
	// TODO
	//
	"bytes"
	"fmt"
	"strconv"
)

type StringBuilder struct {
	buffer *bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	builder1 := new(StringBuilder)
	builder1.buffer = new(bytes.Buffer)
	builder1.buffer.Grow(0xff)

	return builder1
}

// 支持string, Error(), String(), int, int64, float64
func (builder *StringBuilder) Append(a interface{}) {
	var message1 string
	switch a.(type) {
	case string:
		message1 = a.(string)
	case fmt.Stringer:
		message1 = a.(fmt.Stringer).String()
	case error:
		message1 = a.(error).Error()
	case int:
		message1 = strconv.Itoa(a.(int))
	case int64:
		message1 = strconv.FormatInt(a.(int64), 10)
	case float64:
		message1 = strconv.FormatFloat(a.(float64), 'f', -1, 64)
	case bool:
		message1 = strconv.FormatBool(a.(bool))
	default:
		message1 = ""
	}
	if builder.buffer.Cap()-builder.buffer.Len() <= len(message1) {
		builder.buffer.Grow(len(message1) + 0xff)
	}
	builder.buffer.WriteString(message1)
	return
}

func (builder *StringBuilder) String() string {
	return builder.buffer.String()
}
