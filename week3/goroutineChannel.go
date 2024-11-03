// File: goroutineChannel.go
//go:build goroutineChannel
// +build goroutineChannel

package main

import (
	"fmt"
)

func add(a int, b int, outputChan chan int) {
	outputChan <- a + b
	close(outputChan)
}

func multiply(a int, b int, outputChan chan int) {
	outputChan <- a * b
	close(outputChan)
}

func main() {
	var a int
	var b int
	addOutputChan := make(chan int)
	mulOutputChan := make(chan int)
	fmt.Print("첫 번째 정수를 입력하세요: ")
	fmt.Scan(&a)
	fmt.Print("두 번째 정수를 입력하세요: ")
	fmt.Scan(&b)

	go add(a, b, addOutputChan)
	go multiply(a, b, mulOutputChan)

	fmt.Println("덧셈 결과는 :", <-addOutputChan)
	fmt.Println("곱셈 결과는 :", <-mulOutputChan)

}
