// File: atm.go
//go:build atm
// +build atm

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Errorf 함수를 이용하여 에러 메시지를 출력할 수 있습니다.
// %d는 정수, %s는 문자열을 출력할 때 사용합니다. %v는 임의의 타입을 출력할 때 사용합니다.
// fmt.Errorf("에러 메시지 %v", 변수1, 변수2, ...)

type BankAccount struct {
	balance int
}

var minusAmountError error = errors.New("오류: 음수 금액은 입력할 수 없습니다")

var negativeBalanceError error = errors.New("오류: 잔액이 부족합니다")

func (b *BankAccount) Deposit(amount int) error {
	if amount < 0 {
		fmt.Printf("%v : %d원 \n\n", minusAmountError, amount)
		return minusAmountError
	}
	b.balance += amount
	fmt.Printf("입금 성공! 현재 잔액: %v원 \n\n", b.balance)
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount < 0 {
		fmt.Printf("%v : %d원 \n\n", minusAmountError, amount)
		return minusAmountError
	} else if b.balance < amount {
		fmt.Printf("%v : 현재 잔액: %d원, 요청한 금액: %d원 \n\n", negativeBalanceError, b.balance, amount)
		return negativeBalanceError
	}

	b.balance -= amount
	fmt.Printf("출금 성공! 현재 잔액: %v원\n\n", b.balance)
	return nil
}

func main() {
	var account BankAccount

	var input string

	for {
		fmt.Printf("입금 (1), 출금 (2), 종료 (Others) : ")
		fmt.Scan(&input)
		if input == "1" {
			fmt.Printf("입금할 금액을 입력하세요: ")
			fmt.Scan(&input)

			amount, _ := strconv.Atoi(input)
			account.Deposit(amount)
		} else if input == "2" {
			fmt.Printf("출금할 금액을 입력하세요: ")
			fmt.Scan(&input)
			amount, _ := strconv.Atoi(input)
			account.Withdraw(amount)

		} else {
			break
		}
	}
}
