// File: goroutineChannel2.go
//go:build goroutineChannel2
// +build goroutineChannel2

package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	addOutputChan := make(chan int)
	mulOutputChan := make(chan int)
	exitChan := make(chan bool, 2)
	fmt.Print("첫 번째 정수를 입력하세요: ")
	fmt.Scan(&a)
	fmt.Print("두 번째 정수를 입력하세요: ")
	fmt.Scan(&b)

	go func(a int, b int, outputChan chan int, exitChan chan bool) {
		for {
			select {
			case outputChan <- a + b:
			case <-exitChan:
				return
			}
		}
	}(a, b, addOutputChan, exitChan)

	go func(a int, b int, outputChan chan int, exitChan chan bool) {
		for {
			select {
			case outputChan <- a * b:
			case <-exitChan:
				return
			}
		}
	}(a, b, mulOutputChan, exitChan)

	fmt.Println("덧셈 결과는 :", <-addOutputChan)
	fmt.Println("곱셈 결과는 :", <-mulOutputChan)

	exitChan <- true
	exitChan <- true

}
