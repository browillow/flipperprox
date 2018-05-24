package main

import (
	"net/http"
)

// Certbot responds to a Certbot challenge
func Certbot(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/.well-known/acme-challenge/"+Settings.CertbotChallengePrompt {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(Settings.CertbotChallengeResponse))
			return
		}
		next.ServeHTTP(w, r)
	})
}
