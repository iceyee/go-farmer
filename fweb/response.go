package fweb

import (
	"net/http"
	//
)

// OK.
func R200(w http.ResponseWriter) {
	http.Error(w, "OK", 200)
	return
}

// No Content.
func R204(w http.ResponseWriter) {
	http.Error(w, "", 204)
	return
}

// Bad Request.
func R400(w http.ResponseWriter) {
	http.Error(w, "Bad Request", 400)
	return
}

// Unauthorized.
func R401(w http.ResponseWriter) {
	http.Error(w, "Unauthorized", 401)
	return
}

// Forbidden.
func R403(w http.ResponseWriter) {
	http.Error(w, "Forbidden", 403)
	return
}

// Not Found.
func R404(w http.ResponseWriter) {
	http.Error(w, "Not Found", 404)
	return
}

// Method Not Allowed.
func R405(w http.ResponseWriter) {
	http.Error(w, "Method Not Allowed", 405)
	return
}

// Internal Server Error.
func R500(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", 500)
	return
}
