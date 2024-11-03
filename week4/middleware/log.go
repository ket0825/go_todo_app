package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func RequestBodyLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// read request body from bytes stream
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", zap.Error(err))
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		// replace request body with new bytes stream
		// give Close method with io.NopCloser
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}

type rwWrapper struct {
	rw     http.ResponseWriter
	mw     io.Writer
	status int
}

func NewRwWrapper(rw http.ResponseWriter, buf io.Writer) *rwWrapper {
	return &rwWrapper{
		rw: rw,
		mw: io.MultiWriter(rw, buf),
	}
}

func (rw *rwWrapper) Header() http.Header {
	return rw.rw.Header()
}

func (r *rwWrapper) Write(i []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return r.mw.Write(i)
}

func (r *rwWrapper) WriteHeader(statusCode int) {
	r.status = statusCode
	r.rw.WriteHeader(statusCode)
}

func NewLogger(l *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf := &bytes.Buffer{}
			rww := NewRwWrapper(w, buf)
			next.ServeHTTP(rww, r)
			l.Printf("%s", buf)
			l.Printf("%d", rww.status)
		})
	}
}
