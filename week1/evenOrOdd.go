// File: evenOrOdd.go
//go:build evenOrOdd
// +build evenOrOdd

package main

import "fmt"

func main() {
	var num1 int

	fmt.Print("숫자를 입력하세요: ")
	fmt.Scan(&num1)

	if num1%2 == 0 {
		fmt.Println("짝수입니다.")
	} else {
		fmt.Println("홀수입니다.")
	}

}
