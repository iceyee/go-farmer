package web

import (
	"github.com/iceyee/go-farmer/farmer"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	//
)

type FileController struct {
}

func (c *FileController) GetApi() []ApiDocument {
	return []ApiDocument{
		c.a1(),
		c.a2(),
		c.a3(),
	}
}

func (*FileController) a1() ApiDocument {
	return ApiDocument{
		Key:          "file-1-first",
		Url:          "/first",
		Method:       "GET,POST",
		ArgumentType: reflect.TypeOf(t4258{}),
		MapTo:        "First",
		Description:  "读取目录",
		Response:     "成功后返回[]",
	}
}

type t4258 struct {
}

func (*FileController) First(w http.ResponseWriter, r *http.Request, arg interface{}) {
	file1, e := os.Open("/opt/farmer-log/")
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	info1, e := file1.Stat()
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	if !info1.IsDir() {
		http.Error(w, "'/opt/farmer-log/'不是目录", 500)
	}
	info2, e := file1.ReadDir(-1)
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	var result = make([]string, 0, 0xff)
	for _, info3 := range info2 {
		if !info3.IsDir() {
			continue
		}
		result = append(result, info3.Name())
	}
	e = WriteJson(w, &result)
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	return
}

func (*FileController) a2() ApiDocument {
	return ApiDocument{
		Key:          "file-2-second",
		Url:          "/second",
		Method:       "GET,POST",
		ArgumentType: reflect.TypeOf(t1235{}),
		MapTo:        "Second",
		Description:  "读取二级目录",
		Response:     "成功后返回[]",
	}
}

type t1235 struct {
	First string `web:"name:first ###type:string ###require:1"`
}

func (*FileController) Second(w http.ResponseWriter, r *http.Request, arg interface{}) {
	if nil == arg {
		http.Error(w, "", 400)
	}
	var t = arg.(t1235)
	file1, e := os.Open("/opt/farmer-log/" + t.First + "/")
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	info1, e := file1.Stat()
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	if !info1.IsDir() {
		http.Error(w, "'/opt/farmer-log/xxx'不是目录", 500)
	}
	info2, e := file1.ReadDir(-1)
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	var result = make([]string, 0, 0xff)
	for _, info3 := range info2 {
		if info3.IsDir() {
			continue
		}
		result = append(result, info3.Name())
	}
	e = WriteJson(w, &result)
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	return
}

func (*FileController) a3() ApiDocument {
	return ApiDocument{
		Key:          "file-3-third",
		Url:          "/third",
		Method:       "GET,POST",
		ArgumentType: reflect.TypeOf(t8235{}),
		MapTo:        "Third",
		Description:  "读取文件",
		Response:     "成功后返回文件内容",
	}
}

type t8235 struct {
	First  string `web:"name:first ###type:string ###require:1"`
	Second string `web:"name:second ###type:string ###require:1"`
}

func (*FileController) Third(w http.ResponseWriter, r *http.Request, arg interface{}) {
	if nil == arg {
		http.Error(w, "", 400)
	}
	var t = arg.(t8235)
	content1, e := os.ReadFile("/opt/farmer-log/" + t.First + "/" + t.Second)
	if nil != e {
		http.Error(w, e.Error(), 500)
	}
	w.Write(content1)
	return
}

func test2(t *testing.T) {
	ControllerRegistryA.Registry(new(FileController))
    // Listen(":12000")
	var server = httptest.NewServer(new(Server))
	defer server.Close()

	var http1 = farmer.NewHttp()
	http1.SetUrl(server.URL + "/first")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())

	http1 = farmer.NewHttp()
	http1.SetUrl(server.URL + "/second?first=TEST")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())

	http1 = farmer.NewHttp()
	http1.SetUrl(server.URL + "/third?first=TEST&second=TEST-ERROR.log")
	http1.SetVerbose(true)
	farmer.Assert(http1.Request())
	farmer.Assert(200 == http1.GetStatusCode())
	return
}
