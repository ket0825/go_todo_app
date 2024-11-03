package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func MyMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := time.Now()
		h.ServeHTTP(w, r)
		d := int64(time.Since(s).Milliseconds())
		fmt.Printf("end %s(%d ms)\n", s.Format(time.RFC3339), d)
	})
}

type AppVersion string

func VersionAdder(v AppVersion) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("App-Version", string(v))
			w.Header().Add("App-Version", string(v))
			next.ServeHTTP(w, r)
		})
	}
}
