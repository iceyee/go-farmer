package fweb

import (
	"net/http"
	"strings"
	//
)

func init() {
	RegistryInterceptor(new(domainInterceptor))
	AuthorizeDomain("127.0.0.1")
	AuthorizeDomain("0.0.0.0")
	AuthorizeDomain("localhost")
	return
}

var domains []string

type domainInterceptor struct {
}

func (d *domainInterceptor) Filter(path string) bool {
	return true
}

func (d *domainInterceptor) Process(
	session *Session,
	w http.ResponseWriter,
	r *http.Request) bool {

	r.Host = strings.ToLower(r.Host)
	for _, x := range domains {
		if strings.Contains(r.Host, x) {
			return true
		}
	}
	http.Error(w, "Forbidden", 403)
	return false
}

// 域名授权, 只能通过授权域名来访问web.
func AuthorizeDomain(domain string) {
	domains = append(domains, domain)
	return
}
