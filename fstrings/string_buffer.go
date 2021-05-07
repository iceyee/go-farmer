package fstrings

import (
	"bytes"
	"github.com/iceyee/go-farmer/v3/ftype"
	"strconv"
	//
)

// 用于连接字符串.
type StringBuffer struct {
	buffer *bytes.Buffer
}

func NewStringBuffer() *StringBuffer {
	var s *StringBuffer
	s = new(StringBuffer)
	s.buffer = new(bytes.Buffer)
	s.buffer.Grow(0xfff)
	return s
}

// 支持所有基础类型, []byte
func (s *StringBuffer) Append(a interface{}) {
	var a001 string
	switch a.(type) {
	case string:
		a001 = a.(string)
	case ftype.Stringer:
		a001 = a.(ftype.Stringer).String()
	case error:
		a001 = a.(error).Error()
	case bool:
		a001 = strconv.FormatBool(a.(bool))
	case int:
		a001 = strconv.Itoa(a.(int))
	case int64:
		a001 = strconv.FormatInt(a.(int64), 10)
	case int32:
		a001 = strconv.FormatInt(int64(a.(int32)), 10)
	case int16:
		a001 = strconv.FormatInt(int64(a.(int16)), 10)
	case int8:
		a001 = strconv.FormatInt(int64(a.(int8)), 10)
	case uint64:
		a001 = strconv.FormatInt(int64(a.(uint64)), 10)
	case uint32:
		a001 = strconv.FormatInt(int64(a.(uint32)), 10)
	case uint16:
		a001 = strconv.FormatInt(int64(a.(uint16)), 10)
	case uint8:
		a001 = strconv.FormatInt(int64(a.(uint8)), 10)
	case float64:
		a001 = strconv.FormatFloat(a.(float64), 'f', -1, 64)
	case float32:
		a001 = strconv.FormatFloat(float64(a.(float32)), 'f', -1, 32)
	case []byte:
		a001 = string(a.([]byte))
	default:
		a001 = ""
	}
	if s.buffer.Cap() <= s.buffer.Len()+len(a001) {
		s.buffer.Grow(len(a001) + 0xfff)
	}
	s.buffer.WriteString(a001)
	return
}

func (s *StringBuffer) String() string {
	return s.buffer.String()
}
