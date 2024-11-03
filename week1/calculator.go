// File: calculator.go
//go:build calculator
// +build calculator

package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var op string
	fmt.Println("첫 번째 숫자 입력:")
	fmt.Scan(&num1)
	fmt.Println("두 번째 숫자 입력:")
	fmt.Scan(&num2)
	fmt.Println("연산자 입력(+,-,*,/):")
	fmt.Scan(&op)

	if num2 == 0 && op == "/" {
		fmt.Println("에러: 0으로 나눌 수 없습니다.")
	} else {
		switch op {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			fmt.Println(num1 / num2)
		default:
			fmt.Println("잘못된 연산자입니다.")
		}
	}
}
