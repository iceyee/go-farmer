package webframework

import (
	// TODO
	//
	"github.com/iceyee/go-farmer/farmer"
	"net/http"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := r.ParseForm()
	if nil != e {
		http.Error(w, farmer.NewFarmerError(e).Error(), 500)
		return
	}
	if InterceptorRegistryA.Process(w, r) &&
		ControllerRegistryA.Process(w, r) &&
		FileServerA.Process(w, r) {
		http.NotFound(w, r)
		return
	}
	return
}
