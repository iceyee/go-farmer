package fstrings

import (
	"bytes"
	"reflect"
	"strconv"
	//
)

func bufferAppend(b *bytes.Buffer, content string) {
	if b.Cap() <= b.Len()+len(content) {
		b.Grow(0xff + len(content))
	}
	b.WriteString(content)
	return
}

var floatType101 reflect.Type = reflect.TypeOf(float64(0))
var intType101 reflect.Type = reflect.TypeOf(int64(0))
var uintType101 reflect.Type = reflect.TypeOf(uint64(0))

// 把对象转成字符串.
func ObjectToString(o interface{}) string {
	if nil == o {
		return ""
	}
	var buffer *bytes.Buffer
	buffer = new(bytes.Buffer)
	buffer.Grow(0xff)
	objectToString(buffer, reflect.ValueOf(o))
	return buffer.String()
}

func objectToString(b *bytes.Buffer, v reflect.Value) {
	if !v.IsValid() {
		return
	}
	var kind001 reflect.Kind
	kind001 = v.Kind()
	switch kind001 {

	case reflect.Bool:
		boolToString(b, v)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intToString(b, v)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		uintToString(b, v)

	case reflect.Float32, reflect.Float64:
		floatToString(b, v)

	case reflect.Array, reflect.Slice:
		arrayToString(b, v)

	case reflect.Chan, reflect.Func, reflect.Interface:
		pointerToString(b, v)

	case reflect.Map:
		mapToString(b, v)

	case reflect.Ptr:
		ptrToString(b, v)

	case reflect.String:
		stringToString(b, v)

	case reflect.Struct:
		structToString(b, v)

	default:
		return
	}
	return
}

func boolToString(b *bytes.Buffer, v reflect.Value) {
	var a001 string
	if v.Bool() {
		a001 = "true"
	} else {
		a001 = "false"
	}
	bufferAppend(b, a001)
	return
}

func intToString(b *bytes.Buffer, v reflect.Value) {
	var a001 string
	a001 = strconv.FormatInt(v.Int(), 10)
	bufferAppend(b, a001)
	return
}

func uintToString(b *bytes.Buffer, v reflect.Value) {
	var a001 string
	a001 = strconv.FormatInt(int64(v.Uint()), 10)
	bufferAppend(b, a001)
	return
}

func floatToString(b *bytes.Buffer, v reflect.Value) {
	var a001 string
	a001 = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	bufferAppend(b, a001)
	return
}

func arrayToString(b *bytes.Buffer, v reflect.Value) {
	if v.IsNil() {
		bufferAppend(b, "nil")
		return
	}
	bufferAppend(b, "[Array] [")
	var length int
	length = v.Len()
	for x := 0; x < length; x++ {
		var v2 reflect.Value
		v2 = v.Index(x)
		bufferAppend(b, " ")
		objectToString(b, v2)
		bufferAppend(b, ",")
	}
	bufferAppend(b, " ]")
	return
}

func pointerToString(b *bytes.Buffer, v reflect.Value) {
	if v.IsNil() {
		bufferAppend(b, "nil")
		return
	}
	var a001 string
	a001 = strconv.FormatInt(int64(v.Pointer()), 10)
	bufferAppend(b, a001)
	return
}

func mapToString(b *bytes.Buffer, v reflect.Value) {
	if v.IsNil() {
		bufferAppend(b, "nil")
		return
	}
	bufferAppend(b, "[Map] {")
	var keys []reflect.Value
	keys = v.MapKeys()
	for _, key := range keys {
		var v2 reflect.Value
		v2 = v.MapIndex(key)
		if b.Cap() <= b.Len()+0xf {
			b.Grow(0xff)
		}
		bufferAppend(b, " $")
		objectToString(b, key)
		bufferAppend(b, "=")
		objectToString(b, v2)
		bufferAppend(b, ",")
	}
	bufferAppend(b, " }")
	return
}

func ptrToString(b *bytes.Buffer, v reflect.Value) {
	if v.IsNil() {
		bufferAppend(b, "nil")
		return
	}
	bufferAppend(b, v.Type().String())
	bufferAppend(b, " ")
	var a001 string
	a001 = strconv.FormatInt(int64(v.Pointer()), 10)
	bufferAppend(b, a001)
	bufferAppend(b, " ")
	objectToString(b, v.Elem())
	return
}

func stringToString(b *bytes.Buffer, v reflect.Value) {
	bufferAppend(b, v.String())
	return
}

func structToString(b *bytes.Buffer, v reflect.Value) {
	if name := v.Type().Name(); "" == name {
		bufferAppend(b, "[Anonymous] {")
	} else {
		bufferAppend(b, name)
		bufferAppend(b, " {")
	}
	var t reflect.Type
	t = v.Type()
	for x := 0; x < t.NumField(); x++ {
		field := t.Field(x)
		if field.Anonymous {
			continue
		}
		var v2 reflect.Value
		v2 = v.FieldByName(field.Name)
		bufferAppend(b, " $")
		bufferAppend(b, field.Name)
		bufferAppend(b, "=")
		objectToString(b, v2)
		bufferAppend(b, ",")
	}
	bufferAppend(b, " }")
	return
}
