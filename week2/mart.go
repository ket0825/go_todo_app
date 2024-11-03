// File: mart.go
//go:build mart
// +build mart

package main

import "fmt"

type Snack struct {
	Name  string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (s Snack) SaleAndGetPrice() int {
	return int(s.Price * 0.9)
}

func (d Drink) SaleAndGetPrice() int {
	return int(d.Price * 0.8)
}

type Item interface {
	SaleAndGetPrice() int
}

func SaleAndGetPrice(item Item) int {
	return item.SaleAndGetPrice()
}

func main() {
	chips := Snack{
		Name:  "Pringles",
		Price: 4000,
	}
	cracker := Snack{
		Name:  "Ace",
		Price: 2500,
	}

	soda := Drink{
		Name:  "Sprite",
		Price: 1800,
	}
	coffee := Drink{
		Name:  "Top",
		Price: 2700,
	}

	total := 0
	total += SaleAndGetPrice(chips)
	total += SaleAndGetPrice(cracker)
	total += SaleAndGetPrice(soda)
	total += SaleAndGetPrice(coffee)

	fmt.Println(total)
}
