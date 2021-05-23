package fweb

import (
	"github.com/iceyee/go-farmer/v4/flog"
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
		R500(w)
		return
	}
	var session *Session
	if cookie, e := r.Cookie("session_id"); nil != cookie && nil == e {
		var ok bool
		session, ok = getSession(cookie.Value)
		if !ok {
			w.Header().Add("Set-Cookie", "session_id="+session.id+"; Path=/")
		}
	} else {
		session, _ = getSession("")
		w.Header().Add("Set-Cookie", "session_id="+session.id+"; Path=/")
	}
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
		var Banner string = `<pre>
  _________         __         _______       ____    ____    _________    _______    
 |_   ___  |       /  \       |_   __ \     |_   \  /   _|  |_   ___  |  |_   __ \  
   | |_  \_|      / /\ \        | |__) |      |   \/   |      | |_  \_|    | |__) | 
   |  _|         / ____ \       |  __ /       | |\  /| |      |  _|  _     |  __ /  
  _| |_        _/ /    \ \_    _| |  \ \_    _| |_\/_| |_    _| |___/ |   _| |  \ \_
 |_____|      |____|  |____|  |____| |___|  |_____||_____|  |_________|  |____| |___|
</pre>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(Banner))
		return
	} else {
		R404(w)
		return
	}
}
