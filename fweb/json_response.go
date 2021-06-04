package fweb

import (
	"net/http"
	//
)

// OK.
func J200(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "OK.",
		Result:  true,
	}
	WriteJson(w, data)
	return
}

// No Content.
func J204(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "No Content.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}

// Bad Request.
func J400(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "Bad Request.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}

// Unauthorized.
func J401(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "Unauthorized.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}

// Forbidden.
func J403(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "Forbidden.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}

// Not Found.
func J404(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "Not Found.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}

// Method Not Allowed.
func J405(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "Method Not Allowed.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}

// Internal Server Error.
func J500(w http.ResponseWriter) {
	var data = T404{
		Data:    nil,
		Message: "Internal Server Error.",
		Result:  false,
	}
	WriteJson(w, data)
	return
}
