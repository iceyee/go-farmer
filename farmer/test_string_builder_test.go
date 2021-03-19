package farmer

import (
	// TODO
	//
	"testing"
)

func TestStringBuilder(t *testing.T) {
	builder1 := NewStringBuilder()
	builder1.Append("hello ")
	builder1.Append("world ")
	builder1.Append("!!!")
	string1 := builder1.String()
	println(string1)
	Assert("hello world !!!" == string1)
	return
}
