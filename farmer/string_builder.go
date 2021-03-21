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
	b := new(StringBuilder)
	b.buffer = new(bytes.Buffer)
	b.buffer.Grow(0xff)

	return b
}

// 支持string, Error(), String(), int, int64, float64, bool
func (b *StringBuilder) Append(a interface{}) {
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
	if b.buffer.Cap()-b.buffer.Len() <= len(message1) {
		b.buffer.Grow(len(message1) + 0xff)
	}
	b.buffer.WriteString(message1)
	return
}

func (b *StringBuilder) String() string {
	return b.buffer.String()
}
