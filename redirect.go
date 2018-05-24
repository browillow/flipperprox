package main

import (
	"net/http"
	"strings"
)

func getPathBase(path string) string {
	parts := strings.Split(path, "/")
	return parts[1]
}

func getStandardBase(path string) string {
	base := getPathBase(path)
	standardBases := []string{"app", "api"}
	for _, standardBase := range standardBases {
		if base == standardBase {
			return standardBase
		}
	}
	return ""
}

// RedirectCheck causes a redirect from HTTP to HTTPS and from an unsupported base path to /app/...
func RedirectCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		//isHTTP := r.Header.Get("x-forwarded-proto") == "http"
		base := getStandardBase(path)
		if base == "" {
			pathPrepend := ""
			if base == "" {
				pathPrepend = "/app"
			}
			redirectURL := "http://" + r.Host + pathPrepend + r.URL.Path
			if len(r.URL.RawQuery) > 0 {
				redirectURL += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, redirectURL, http.StatusPermanentRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}
