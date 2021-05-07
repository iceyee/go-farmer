package fhttp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/flog"
	"github.com/iceyee/go-farmer/v3/fstrings"
	"github.com/iceyee/go-farmer/v3/ftype"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	//
)

type Http struct {
	body    []byte            // POST的内容
	e       error             // 异常, 在Request()返回false后, 可调用GetError()获得异常结果
	header  map[string]string // 请求头
	method  string            // 请求方法
	proxy   string            // 代理
	timeout int64             // 超时, 默认60秒
	url     string            // 链接

	responseBody   []byte              // 响应内容
	responseHeader map[string][]string // 响应头
	statusCode     int                 // 状态码
}

func init() {
	rand.Seed(time.Now().UnixNano())
	return
}

func New() *Http {
	h := new(Http)
	h.method = "GET"
	h.timeout = 1000 * 60
	return h
}

func (h *Http) SetBody(body []byte) {
	h.body = body
	return
}

func (h *Http) SetHeader(header map[string]string) {
	h.header = header
	return
}

func (h *Http) SetMethod(method string) {
	h.method = method
	return
}

func (h *Http) SetTimeout(timeout int64) {
	if 0 < timeout {
		h.timeout = timeout
	}
	return
}

// @proxy: "", 或完整的代理, 如"socks5://[user:password@]hk.farmer.ink:10002"
func (h *Http) SetProxy(proxy string) {
	h.proxy = proxy
	return
}

func (h *Http) SetUrl(url string) {
	h.url = url
	return
}

func (h *Http) GetError() ftype.Error {
	return h.e
}

func (h *Http) GetStatusCode() int {
	return h.statusCode
}

func (h *Http) GetResponseHeader() map[string][]string {
	return h.responseHeader
}

func (h *Http) GetResponseBody() []byte {
	return h.responseBody
}

// 发起请求, 其中URL是必须填的, 如果返回false, 则表示发生异常, 可以调用.GetError()获取异常.
// 只有返回true的时候, .GetStatusCode(), .GetResponseHeader(), .GetResponseBody() 才会有效
func (h *Http) Request() bool {
	var sb001 *fstrings.StringBuffer
	sb001 = fstrings.NewStringBuffer()
	defer flog.Debug(sb001)
	sb001.Append("fhttp.Http - ")
	sb001.Append("\nUrl: ")
	sb001.Append(h.url)
	sb001.Append("\nMethod: ")
	sb001.Append(h.method)
	sb001.Append("\nHeader: ")
	sb001.Append(fmt.Sprintf("%v", h.header))
	sb001.Append("\nTimeout: ")
	sb001.Append(h.timeout)
	sb001.Append("\nProxy: ")
	sb001.Append(h.proxy)
	sb001.Append("\n")
	if "" == h.url ||
		"POST" == h.method && nil == h.body {

		h.e = ferror.New("参数不完整, 缺少url或body")
		return false
	}
	var body001 io.Reader
	if nil != h.body {
		body001 = bytes.NewReader(h.body)
	}
	var request001 *http.Request
	var e error
	request001, e = http.NewRequest(h.method, h.url, body001)
	if nil != e {
		h.e = ferror.New(e)
		return false
	}
	request001.Close = true
	if nil != h.header {
		for key, value := range h.header {
			request001.Header.Set(key, value)
		}
	}
	if _, ok := request001.Header["Accept-Encoding"]; !ok {
		request001.Header.Set("Accept-Encoding", "gzip")
	}
	if _, ok := request001.Header["Accept-Language"]; !ok {
		request001.Header.Set("Accept-Language", "zh, zh-CN")
	}
	if _, ok := request001.Header["User-Agent"]; !ok {
		request001.Header.Set("User-Agent", "iceyee/2.0")
	}
	if _, ok := request001.Header["X-Requested-With"]; !ok {
		request001.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	if _, ok := request001.Header["X-Forwarded-For"]; !ok {
		// ip1, ip2, ip3 - IP, 临时变量
		a := [8]byte{}
		for index, _ := range a {
			a[index] = byte(rand.Intn(0xff))
		}
		ip1 := net.IPv4(a[0], a[1], a[2], a[3]).String()
		ip2 := net.IPv4(a[4], a[5], a[6], a[7]).String()
		ip3 := fmt.Sprintf("%s, %s", ip1, ip2)
		request001.Header.Set("X-Forwarded-For", ip3)
	}
	if nil != h.body {
		request001.Header.Set(
			"Content-Length",
			strconv.FormatInt(int64(len(h.body)), 10))
		if _, ok := request001.Header["Content-Type"]; !ok {
			request001.Header.Set(
				"Content-Type",
				"application/x-www-form-urlencoded; charset=UTF-8")
		}
	}
	sb001.Append("\n>>>>>>>>>>\n")
	sb001.Append(h.method)
	sb001.Append(" ")
	sb001.Append(h.url)
	for key, values := range request001.Header {
		for _, value := range values {
			sb001.Append("\n")
			sb001.Append(key)
			sb001.Append(": ")
			sb001.Append(value)
		}
	}
	sb001.Append("\n")
	if nil != h.body {
		sb001.Append("\n")
		sb001.Append(string(h.body))
		sb001.Append("\n")
	}
	var client001 *http.Client
	client001 = new(http.Client)
	client001.Timeout = time.Duration(h.timeout) * time.Millisecond
	if "" != h.proxy {
		var transport001 *http.Transport
		transport001 = new(http.Transport)
		transport001.Proxy = func(*http.Request) (*url.URL, error) {
			return url.Parse(h.proxy)
		}
		client001.Transport = transport001
	}
	var response001 *http.Response
	response001, e = client001.Do(request001)
	if nil != e {
		h.e = ferror.New(e)
		return false
	}
	// data - 响应的内容([]byte)
	// encodings1 - 响应头的Content-Encoding([]string)
	// encoding1 - 响应头的Content-Encoding(string)
	// gzipReader001 - (*gzip.Reader)
	defer response001.Body.Close()
	var data []byte
	encodings1, ok := response001.Header["Content-Encoding"]
	if ok {
		encoding1 := strings.Join(encodings1, ",")
		if strings.Contains(strings.ToLower(encoding1), "gzip") {
			var reader001 *gzip.Reader
			reader001, e = gzip.NewReader(response001.Body)
			if nil != e {
				h.e = ferror.New(e)
				return false
			}
			defer reader001.Close()
			data, e = ioutil.ReadAll(reader001)
			if nil != e {
				h.e = ferror.New(e)
				return false
			}
		}
	}
	if nil == data {
		data, e = ioutil.ReadAll(response001.Body)
		if nil != e {
			h.e = ferror.New(e)
			return false
		}
	}
	sb001.Append("\n<<<<<<<<<<\n")
	sb001.Append(response001.Status)
	sb001.Append(" ")
	sb001.Append(response001.Proto)
	for key, values := range response001.Header {
		for _, value := range values {
			sb001.Append("\n")
			sb001.Append(key)
			sb001.Append(": ")
			sb001.Append(value)
		}
	}
	sb001.Append("\n\n")
	sb001.Append(" ")
	sb001.Append(string(data))
	h.statusCode = response001.StatusCode
	h.responseHeader = response001.Header
	h.responseBody = data
	return true
}
