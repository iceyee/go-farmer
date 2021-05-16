package fweb

import (
	"github.com/iceyee/go-farmer/v4/flog"
	"math/rand"
	"time"
	//
)

type t383 struct {
	Context map[string]string
	Time    int64
}

// 保存全部的会话信息.
var context map[string]t383 = make(map[string]t383, 0xffff)

// 当前时间戳, 一秒刷新一次.
var time515 int64 = time.Now().Unix()

func init() {
	rand.Seed(time.Now().UnixNano())
	// 刷新时间戳.
	go func() {
		for true {
			time.Sleep(1 * time.Second)
			time515 = time.Now().Unix()
		}
	}()
	// 删除过期的会话.
	go func() {
		var oneHour int64
		oneHour = 1 * 60 * 60
		for true {
			var a001 []string
			a001 = make([]string, 0, 0xffff)
			for key, value := range context {
				if oneHour < time515-value.Time {
					a001 = append(a001, key)
				}
			}
			for _, x := range a001 {
				delete(context, x)
			}
			time.Sleep(1 * time.Minute)
		}
	}()
	return
}

// 请求会话.
type Session struct {
	id      string
	storage map[string]string
}

// 取得会话信息, 如果不存在则创建.
func getSession(sessionId string) (*Session, bool) {
	var s *Session
	s = new(Session)
	if a001, ok := context[sessionId]; ok {
		s.id = sessionId
		s.storage = a001.Context
		a001.Time = time515
		context[s.id] = a001
		return s, true
	} else {
		var a002 [16]byte
		for x := 0; x < 16; x++ {
			a002[x] = byte(rand.Intn(26)) + byte('a')
		}
		s.id = string(a002[:])
		s.storage = make(map[string]string, 0xff)
		context[s.id] = t383{
			Context: s.storage,
			Time:    time515,
		}
		return s, false
	}
}

func (s *Session) Get(key string) string {
	if nil == s.storage {
		flog.Debug("无效的会话. *fweb.Session")
		return ""
	}
	return s.storage[key]
}

func (s *Session) Set(key string, value string) {
	if nil == s.storage {
		flog.Debug("无效的会话. *fweb.Session")
		return
	}
	s.storage[key] = value
	return
}
