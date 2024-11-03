package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"week4/middleware"
)

func main() {
	// init handler
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	}
	var handler http.Handler = http.HandlerFunc(hello)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	setLogger := middleware.NewLogger(logger)

	const Version = "1.0.0"
	// add recovery
	handler = middleware.RecoveryMiddleware(handler)
	// add logger
	handler = setLogger(handler)
	// add middleware
	handler = middleware.MyMiddleware(handler)
	// add version
	addVersion := middleware.VersionAdder(Version)
	// add version to handler
	handler = addVersion(handler)
	// add log
	handler = middleware.RequestBodyLogMiddleware(handler)

	err := http.ListenAndServe(":18080", handler)

	if err != nil {
		fmt.Printf("Failed to terminate server: %v", err)
		os.Exit(1)
	}

}
