package main

import (
	"net/http"
)

// ProxyRouter routes the request to the reverse proxy handler
func ProxyRouter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathBase := getPathBase(r.URL.Path)
		switch pathBase {
		case "app":
			ProxyHandlers[FlipperApp].ServeHTTP(w, r)
		case "api":
			ProxyHandlers[FlipperAPI].ServeHTTP(w, r)
		default:
			next.ServeHTTP(w, r)
		}
		return
	})
}
