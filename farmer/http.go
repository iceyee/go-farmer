package farmer

import (
	// TODO
	//
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// http://api.k780.com/?app=ip.local&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json
type Http struct {
	body    []byte
	e       error
	header  map[string]string
	method  string
	proxy   string
	url     string
	verbose bool

	responseBody   []byte
	responseHeader map[string][]string
	statusCode     int
}

func init() {
	rand.Seed(time.Now().UnixNano())
	return
}

func NewHttp() *Http {
	http1 := new(Http)
	http1.method = "GET"
	http1.verbose = false
	return http1
}

func (http1 *Http) SetBody(body []byte) {
	http1.body = body
	return
}

func (http1 *Http) SetHeader(header map[string]string) {
	http1.header = header
	return
}

func (http1 *Http) SetMethod(method string) {
	http1.method = method
	return
}

// @proxy: "", 或完整的代理, 如"socks://[user:password@]hk.farmer.ink:10002"
func (http1 *Http) SetProxy(proxy string) {
	http1.proxy = proxy
	return
}

func (http1 *Http) SetUrl(url string) {
	http1.url = url
	return
}

func (http1 *Http) SetVerbose(verbose bool) {
	http1.verbose = verbose
	return
}

func (http1 *Http) GetError() error {
	return http1.e
}

func (http1 *Http) GetStatusCode() int {
	return http1.statusCode
}

func (http1 *Http) GetResponseHeader() map[string][]string {
	return http1.responseHeader
}

func (http1 *Http) GetResponseBody() []byte {
	return http1.responseBody
}

// 发起请求, 其中URL是必须填的, 如果返回false, 则表示发生异常, 可以调用.GetError()获取异常.
// 只有返回true的时候, .GetStatusCode(), .GetResponseHeader(), .GetResponseBody() 才会有效
func (http1 *Http) Request() bool {
	if "" == http1.url || "POST" == http1.method && nil == http1.body {
		http1.e = NewFarmerError("参数不完整, 缺少url或body")
		return false
	}
	var body1 io.Reader
	if "POST" == http1.method && nil != http1.body {
		body1 = bytes.NewReader(http1.body)
	}
	request1, e := http.NewRequest(http1.method, http1.url, body1)
	if nil != e {
		http1.e = NewFarmerError(e)
		return false
	}
	request1.Close = true
	if nil != http1.header {
		for key, value := range http1.header {
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
        a := [8]byte{}
        for index, _ := range a {
            a[index] = byte(rand.Intn(256))
        }
        ip1 := net.IPv4(a[0], a[1], a[2], a[3]).String()
        ip2 := net.IPv4(a[4], a[5], a[6], a[7]).String()
        ip3 := fmt.Sprintf("%s, %s", ip1, ip2)
        request1.Header.Set("X-Forwarded-For", ip3)
    }
	if http1.verbose {
		println(">>>\n", http1.method, http1.url)
		for key, values := range request1.Header {
			for _, value := range values {
				println(key, ":", value)
			}
		}
		println()
		if nil != http1.body {
			println(string(http1.body))
		}
	}

	client1 := new(http.Client)
	client1.Timeout = 60 * time.Second
	if "" != http1.proxy {
		transport1 := new(http.Transport)
		transport1.Proxy = func(*http.Request) (*url.URL, error) {
			return url.Parse(http1.proxy)
		}
		client1.Transport = transport1
	}
	response1, e := client1.Do(request1)

	var data1 []byte
	if nil != e {
		http1.e = NewFarmerError(e)
		return false
	} else {
		defer response1.Body.Close()
		encodings1, exists := response1.Header["Content-Encoding"]
		if exists {
			encoding1 := strings.Join(encodings1, ",")
			if strings.Contains(strings.ToLower(encoding1), "gzip") {
				gzipReader1, e := gzip.NewReader(response1.Body)
				if nil != e {
					http1.e = NewFarmerError(e)
					return false
				}
				defer gzipReader1.Close()

				data1, e = ioutil.ReadAll(gzipReader1)
				if nil != e {
					http1.e = NewFarmerError(e)
					return false
				}
			}
		} else {
			data1, e = ioutil.ReadAll(response1.Body)
			if nil != e {
				http1.e = NewFarmerError(e)
				return false
			}
		}
		if http1.verbose {
			println("\n<<<\n", response1.Status, response1.Proto)
			for key, values := range response1.Header {
				for _, value := range values {
					println(key, ":", value)
				}
			}
			println()
			println(string(data1))
		}
		http1.statusCode = response1.StatusCode
		http1.responseHeader = response1.Header
		http1.responseBody = data1
		return true
	}
}
