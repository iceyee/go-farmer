package fweb

import (
	"github.com/iceyee/go-farmer/v3/ferror"
	"github.com/iceyee/go-farmer/v3/flog"
	"net/http"
	"strings"
	//
)

// 入口. 用httptest时会用到.
type Server struct {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flog.Debug(r.URL.Path + ", " + r.URL.RawQuery)
	var e error
	e = r.ParseForm()
	if nil != e {
		http.Error(w, ferror.New(e).Error(), 500)
		return
	}
	var session *Session
	if cookie, e := r.Cookie("session_id"); nil != cookie && nil == e {
		session = getSession(cookie.Value)
	} else {
		session = getSession("")
	}
	w.Header().Add("Set-Cookie", "session_id="+session.id+"; Path=/")
	for _, interceptor := range x494 {
		if interceptor.Filter(r.URL.Path) {
			if !interceptor.Process(session, w, r) {
				return
			}
		}
	}
	if controller, ok := x251[r.URL.Path]; ok {
		processController(session, w, r, controller)
		return
	}
	for _, fileserver := range x766 {
		if strings.HasPrefix(r.URL.Path, fileserver.Prefix) {
			fileserver.Handler.ServeHTTP(w, r)
			return
		}
	}
	if "/" == r.URL.Path {
		content := `<h3 style="padding: 1em 0; text-align: center;">Author: Farmer</h3>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(content))
		return
	} else {
		http.NotFound(w, r)
		return
	}
}
