package web

import (
	"github.com/iceyee/go-farmer/v2/farmer"
	"net/http"
	//
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if nil != e {
		http.Error(w, farmer.NewFarmerError(e).Error(), 500)
		return
	}
	if InterceptorRegistryA.process(w, r) &&
		ControllerRegistryA.process(w, r) &&
		FileServerA.process(w, r) {
		if "/" == r.URL.Path {
			content1 := `<h3 style="padding: 1em 0; text-align: center;">Author: Farmer</h3>`
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(content1))
			return
		} else {
			http.NotFound(w, r)
			return
		}
	}
	return
}
