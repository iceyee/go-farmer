package futil

import (
	"github.com/iceyee/go-farmer/v4/fassert"
	"log"
	"testing"
	"time"
	//
)

func TestBanner(t *testing.T) {
    t.Log(Banner)
	return
}

func TestDeleteExpiredKey(t *testing.T) {
	var a001 map[string]int64
	a001 = make(map[string]int64, 0xff)
	a001["a"] = 0
	a001["b"] = time.Now().Unix() - 10
	a001["c"] = time.Now().Unix() - 5
	a001["d"] = time.Now().Unix()
	t.Log(a001)
	fassert.Assert(4 == len(a001), "map初始元素是4个")
	DeleteExpiredKey(a001, time.Now().Unix()-6)
	t.Log(a001)
	fassert.Assert(2 == len(a001), "删除6秒前的键, 还剩两个")
	return
}

func TestDeleteExpiredKey2(t *testing.T) {
	var a001 map[string]T737
	a001 = make(map[string]T737, 0xff)
	a001["a"] = T737{Time: 0}
	a001["b"] = T737{Time: time.Now().Unix() - 10}
	a001["c"] = T737{Time: time.Now().Unix() - 5}
	a001["d"] = T737{Time: time.Now().Unix()}
	t.Log(a001)
	fassert.Assert(4 == len(a001), "map初始元素是4个")
	DeleteExpiredKey2(a001, time.Now().Unix()-6)
	t.Log(a001)
	fassert.Assert(2 == len(a001), "删除6秒前的键, 还剩两个")
	return
}

func TestSleep(t *testing.T) {
	Sleep(1000)
	return
}

func TestNowTime(t *testing.T) {
	a := NowTime
	time.Sleep(1 * time.Second)
	b := NowTime
	fassert.Assert(a < b, "a < b")
	return
}

func ExampleDeleteExpiredKey() {
	var a001 map[string]int64
	a001 = make(map[string]int64, 0xff)
	a001["a"] = 0
	a001["b"] = time.Now().Unix() - 10
	a001["c"] = time.Now().Unix() - 5
	a001["d"] = time.Now().Unix()
	log.Println(a001)
	fassert.Assert(4 == len(a001), "map初始元素是4个")
	DeleteExpiredKey(a001, time.Now().Unix()-6)
	log.Println(a001)
	fassert.Assert(2 == len(a001), "删除6秒前的键, 还剩两个")
	return
}

func ExampleDeleteExpiredKey2() {
	var a001 map[string]T737
	a001 = make(map[string]T737, 0xff)
	a001["a"] = T737{Time: 0}
	a001["b"] = T737{Time: time.Now().Unix() - 10}
	a001["c"] = T737{Time: time.Now().Unix() - 5}
	a001["d"] = T737{Time: time.Now().Unix()}
	log.Println(a001)
	fassert.Assert(4 == len(a001), "map初始元素是4个")
	DeleteExpiredKey2(a001, time.Now().Unix()-6)
	log.Println(a001)
	fassert.Assert(2 == len(a001), "删除6秒前的键, 还剩两个")
	return
}

func ExampleSleep() {
	// 延时1秒
	Sleep(1000)
	return
}
