package main

import (
	"errors"
	"fmt"
	"net/http"
)

// RecoverFromPanics catches panics that occur during the HTTP handler chain
func RecoverFromPanics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				fmt.Println("SYSTEM ERROR [FLIPPERPROX] Recovered from panic ->", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
