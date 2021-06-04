package fweb

import (
	"github.com/iceyee/go-farmer/v5/fhttp"
	"github.com/iceyee/go-farmer/v5/flog"
	"net/http"
	"net/http/httptest"
	"testing"
	//
)

type A struct {
}

func (*A) Filter(path string) bool {
	return true
}

func (*A) Process(
	session *Session,
	w http.ResponseWriter,
	r *http.Request) bool {

	if "DELETE" == r.Method {
		r.Method = "GET"
	}
	return true
}

type B struct {
}

func (*B) Test__() string {
	return `
       @Url /test

       @MapTo Test

       @Method GET

       @Description 这是描述

       @Response 响应说明

       @Remarks 备注

       @Parameter | A | int64 | 1 |   | 这是A
       @Constraints | A | 0 | 100 | 50 |

       @Parameter | B | float64 |   | 1 | 这是B
       @Constraints | B | 0 | 100 | 50 |

       @Parameter | C | string | 1 |   | 这是C
       @Constraints | C |   |   | 50 | hello

       @Parameter | D | int64 |   |   | 这是D
       @Constraints | D |   |   |   |
       `
}

type t struct {
	A int64
	B float64
	C string
	D int64
}

func (*B) Test(
	session *Session,
	w http.ResponseWriter,
	r *http.Request,
	A *int64,
	B *float64,
	C *string,
	D *int64) {

	a := t{
		A: *A,
		B: *B,
		C: *C,
	}
	if nil == D {
		a.D = -1
	} else {
		a.D = *D
	}
	WriteJson(w, a)
	return
}

func TestRegistryController(t *testing.T) {
	flog.SetLogLevel(flog.L_DEBUG)
	RegistryInterceptor(new(A))
	RegistryController(new(B))
	RegistryFileServer("/", "/tmp/")
	SetResponseMode(M_JSON)
	var server = httptest.NewServer(new(Server))
	defer server.Close()
	var h *fhttp.Http
	h = fhttp.New()
	h.SetUrl(server.URL + "/test")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/test?A=1&B=2&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/test?A=1&B=50&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/test?A=50&B=2&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/test?A=1&B=2&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/test?A=-21&B=2&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/test?A=1&B=-2&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetMethod("POST")
	h.SetUrl(server.URL + "/test?A=1&B=2&C=hello%20world.")
	h.SetBody([]byte("hello world"))
	h.Request()

	h = fhttp.New()
	h.SetMethod("DELETE")
	h.SetUrl(server.URL + "/test?A=1&B=2&C=hello%20world.")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/0/status")
	h.Request()

	h = fhttp.New()
	h.SetUrl(server.URL + "/0/api")
	h.Request()

	return
}

func test(t *testing.T) {
	RegistryInterceptor(new(A))
	RegistryController(new(B))
	RegistryFileServer("/", "/tmp/")
	panic(Listen(":9999"))
}
