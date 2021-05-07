package fstrings

import (
	"testing"
	//
)

func TestStringBuffer(t *testing.T) {
	var s *StringBuffer
	s = NewStringBuffer()
	s.Append("hello world.\n")
	s.Append(0xfff)
	s.Append("\n")
	s.Append(0.123456)
	s.Append("\n")
	s.Append(false)
	s.Append("\n")
	t.Log(s)
	return
}

func ExampleStringBuffer() {
	var s *StringBuffer
	s = NewStringBuffer()
	s.Append("hello world.\n")
	s.Append(0xfff)
	s.Append("\n")
	s.Append(0.123456)
	s.Append("\n")
	s.Append(false)
	s.Append("\n")
	println(s.String())
	return
}
