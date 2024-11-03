package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				jsonBody, _ := json.Marshal(map[string]string{
					"error": fmt.Sprintf("%v", err),
				})
				// set response header
				w.Header().Set("Content-Type", "application/json")
				// set response status code
				w.WriteHeader(http.StatusInternalServerError)
				// write response body
				w.Write(jsonBody)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
