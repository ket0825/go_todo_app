package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// http.Handler adapter example.
	hello := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	}
	handler := http.HandlerFunc(hello)
	err := http.ListenAndServe(":18080", handler)

	if err != nil {
		fmt.Printf("Failed to terminate server: %v", err)
		os.Exit(1)
	}

}
