package fweb

import (
	"github.com/iceyee/go-farmer/v5/flog"
	"net/http"
	"strings"
	//
)

func init() {
	RegistryInterceptor(new(domainInterceptor))
	AuthorizeDomain("0.0.0.0")
	AuthorizeDomain("127.0.0.1")
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
	if M_PLAIN == mode {
		R403(w)
	} else if M_JSON == mode {
		J403(w)
	}
	return false
}

// 域名授权, 只能通过授权域名来访问web.
func AuthorizeDomain(domain string) {
	flog.Debug("授权域名, " + domain)
	domains = append(domains, domain)
	return
}
