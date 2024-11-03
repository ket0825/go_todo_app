// File: factorial.go
//go:build factorial
// +build factorial

package main

/*
#cgo LDFLAGS: -L. -lcfactorial
int factorial(int n);
*/
import "C"
import "fmt"

func main() {
	n := 5
	fmt.Printf("Factorial of %d is %d\n", n, C.factorial(C.int(n)))
}
