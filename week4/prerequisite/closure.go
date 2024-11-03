//go:build closure
// +build closure

// File: closure.go
package main

import (
	"fmt"
)

func closure() func(int, string) string {
	cache := make(map[int]string)
	return func(n int, s string) string {
		if val, ok := cache[n]; ok {
			return fmt.Sprintf("cached value %d: %s", n, val)
		} else {
			cache[n] = s
			return fmt.Sprintf("not cached: %d", n)
		}
	}
}

func main() {
	// Closure example
	// The closure is a function
	// that captures the environment in which it was created.
	cacheClosure := closure()
	fmt.Println(cacheClosure(1, "one"))
	fmt.Println(cacheClosure(2, "two"))
	fmt.Println(cacheClosure(1, "one"))
}
