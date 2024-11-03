// File: fibonacci.go
//go:build fibonacci
// +build fibonacci

package main

/*
#cgo LDFLAGS: -L. -lcfibonacci
int fibonacci(int n);
*/
import "C"
import "fmt"

func main() {
	n := 10
	fmt.Printf("fibonacci of %d is %d\n", n, C.fibonacci(C.int(n)))
}
