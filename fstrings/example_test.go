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

func TestObjectToString(t *testing.T) {
	var Config struct {
		Redis struct {
			Db       int64
			Host     string
			Password string
			Port     int64
		}
		RabbitMQ struct {
			User     string
			Host     string
			Password string
			Port     int64
		}
		A map[string]int64
		B []string
	}
	t.Log(ObjectToString(Config))
	Config.A = make(map[string]int64, 0xff)
	Config.A["hello"] = 1
	Config.A["world"] = 2
	t.Log(ObjectToString(Config))
	Config.B = make([]string, 0, 0xff)
	Config.B = append(Config.B, "hello")
	Config.B = append(Config.B, "world")
	t.Log(ObjectToString(Config))
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

func ExampleObjectToString() {
	var Config struct {
		Redis struct {
			Db       int64
			Host     string
			Password string
			Port     int64
		}
		RabbitMQ struct {
			User     string
			Host     string
			Password string
			Port     int64
		}
		A map[string]int64
		B []string
	}
	println(ObjectToString(Config))
	Config.A = make(map[string]int64, 0xff)
	Config.A["hello"] = 1
	Config.A["world"] = 2
	println(ObjectToString(Config))
	Config.B = make([]string, 0, 0xff)
	Config.B = append(Config.B, "hello")
	Config.B = append(Config.B, "world")
	println(ObjectToString(Config))
	return
}
