package farmer

import (
	"bytes"
	"compress/gzip"
	"fmt"
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
	url     string            // 链接
	verbose bool              // 冗余模式

	responseBody   []byte              // 响应内容
	responseHeader map[string][]string // 响应头
	statusCode     int                 // 状态码
}

func init() {
	rand.Seed(time.Now().UnixNano())
	return
}

func NewHttp() *Http {
	h := new(Http)
	h.method = "GET"
	h.verbose = false
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

// @proxy: "", 或完整的代理, 如"socks://[user:password@]hk.farmer.ink:10002"
func (h *Http) SetProxy(proxy string) {
	h.proxy = proxy
	return
}

func (h *Http) SetUrl(url string) {
	h.url = url
	return
}

func (h *Http) SetVerbose(verbose bool) {
	h.verbose = verbose
	return
}

func (h *Http) GetError() error {
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
	if "" == h.url || "POST" == h.method && nil == h.body {
		h.e = NewFarmerError("参数不完整, 缺少url或body")
		return false
	}
	// body1 - POST内容(bytes.Reader)
	// request1 - 请求(*http.Request)
	var body1 io.Reader
	if nil != h.body {
		body1 = bytes.NewReader(h.body)
	}
	request1, e := http.NewRequest(h.method, h.url, body1)
	if nil != e {
		h.e = NewFarmerError(e)
		return false
	}
	request1.Close = true
	if nil != h.header {
		for key, value := range h.header {
			request1.Header.Set(key, value)
		}
	}
	if _, exists := request1.Header["Accept-Encoding"]; !exists {
		request1.Header.Set("Accept-Encoding", "gzip")
	}
	if _, exists := request1.Header["Accept-Language"]; !exists {
		request1.Header.Set("Accept-Language", "zh, zh-CN")
	}
	if _, exists := request1.Header["User-Agent"]; !exists {
		request1.Header.Set("User-Agent", "farmer")
	}
	if _, exists := request1.Header["X-Requested-With"]; !exists {
		request1.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	if _, exists := request1.Header["X-Forwarded-For"]; !exists {
		// ip1, ip2, ip3 - IP, 临时变量
		a := [8]byte{}
		for index, _ := range a {
			a[index] = byte(rand.Intn(256))
		}
		ip1 := net.IPv4(a[0], a[1], a[2], a[3]).String()
		ip2 := net.IPv4(a[4], a[5], a[6], a[7]).String()
		ip3 := fmt.Sprintf("%s, %s", ip1, ip2)
		request1.Header.Set("X-Forwarded-For", ip3)
	}
	if nil != h.body {
		request1.Header.Set("Content-Length", strconv.FormatInt(int64(len(h.body)), 10))
		if _, exists := request1.Header["Content-Type"]; !exists {
			request1.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
		}
	}
	if h.verbose {
		println(">>>\n", h.method, h.url)
		for key, values := range request1.Header {
			for _, value := range values {
				println(key, ":", value)
			}
		}
		println()
		if nil != h.body {
			println(string(h.body))
		}
	}

	// client1 - (*http.Client)
	// transport1 - (*http.Transport)
	// response1 - (*http.Response)
	client1 := new(http.Client)
	client1.Timeout = 60 * time.Second
	if "" != h.proxy {
		transport1 := new(http.Transport)
		transport1.Proxy = func(*http.Request) (*url.URL, error) {
			return url.Parse(h.proxy)
		}
		client1.Transport = transport1
	}
	response1, e := client1.Do(request1)
	if nil != e {
		h.e = NewFarmerError(e)
		return false
	} else {
		// 请求成功, 开始处理数据
		// data1 - 响应的内容([]byte)
		// encodings1 - 响应头的Content-Encoding([]string)
		// encoding1 - 响应头的Content-Encoding(string)
		// gzipReader1 - (*gzip.Reader)
		defer response1.Body.Close()
		var data1 []byte
		encodings1, exists := response1.Header["Content-Encoding"]
		if exists {
			encoding1 := strings.Join(encodings1, ",")
			if strings.Contains(strings.ToLower(encoding1), "gzip") {
				gzipReader1, e := gzip.NewReader(response1.Body)
				if nil != e {
					h.e = NewFarmerError(e)
					return false
				}
				defer gzipReader1.Close()
				data1, e = ioutil.ReadAll(gzipReader1)
				if nil != e {
					h.e = NewFarmerError(e)
					return false
				}
			}
		} else {
			data1, e = ioutil.ReadAll(response1.Body)
			if nil != e {
				h.e = NewFarmerError(e)
				return false
			}
		}
		if h.verbose {
			println("\n<<<\n", response1.Status, response1.Proto)
			for key, values := range response1.Header {
				for _, value := range values {
					println(key, ":", value)
				}
			}
			println()
			println(string(data1))
		}
		h.statusCode = response1.StatusCode
		h.responseHeader = response1.Header
		h.responseBody = data1
		return true
	}
}
