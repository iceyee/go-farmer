package fweb

import (
	"github.com/iceyee/go-farmer/v3/fhttp"
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

func (*A) Process(session *Session, w http.ResponseWriter, r *http.Request) bool {
	if "DELETE" == r.Method {
		r.Method = "GET"
	}
	return true
}

type B struct {
}

func (*B) F123() string {
	return `
    @Description 这是描述

    @MapTo Test

    @Method GET

    @Remarks 备注

    @Response 响应说明

    @Url /test

    @Parameter | A | int64 | 1 |  | 这是A
    @Constraints | A | 0 | 100 | 50 |  

    @Parameter | B | float64 |  | 1 | 这是B
    @Constraints | B | 0 | 100 | 50 |  

    @Parameter | C | string | 1 |  | 这是C
    @Constraints | C |  |  | 50 | hello
    `
}

type t struct {
	A int64
	B float64
	C string
}

func (*B) Test(
	session *Session,
	w http.ResponseWriter,
	r *http.Request,
	A int64,
	B float64,
	C string) {

	a := t{
		A: A,
		B: B,
		C: C,
	}
	WriteJson(w, a)
	return
}

func testRegistryController(t *testing.T) {
	RegistryInterceptor(new(A))
	RegistryController(new(B))
	RegistryFileServer("/", "/tmp/")
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

func Test(t *testing.T) {
	RegistryInterceptor(new(A))
	RegistryController(new(B))
	RegistryFileServer("/", "/tmp/")
	Listen(":9999")
}
