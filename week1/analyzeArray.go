// File: analyzeArray.go
//go:build analyzeArray
// +build analyzeArray

package main

import "fmt"

func sumArray(arr []int) {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	fmt.Println(sum)
}

func findMaxMin(arr []int) {
	max := arr[0]
	min := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}

func judgeArrayLen(arr []int) {
	arrLen := len(arr)
	if arrLen < 3 {
		fmt.Println("배열의 길이가 짧습니다.")
	} else if arrLen == 5 {
		fmt.Println("배열의 길이가 적당합니다.")
	} else {
		fmt.Println("배열의 길이가 깁니다.")
	}
}

func main() {
	defer fmt.Println("프로그램이 종료되었습니다.")
	arr := []int{3, 5, 1, 2, 0}

	sumArray(arr)
	findMaxMin(arr)
	judgeArrayLen(arr)

}
