// File: sort.go
//go:build sort
// +build sort

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func sorting(userList []User) {
	// 1,5,2,4,3
	// 1,2,5,4,3
	// 1,2,4,5,3
	// 1,2,4,3,5
	// 1,2,3,4
	var userListLength = len(userList)
	for i := 0; i < userListLength-1; i++ {
		for j := i; j < userListLength; j++ {
			if userList[i].Age > userList[j].Age {
				change(&userList[i], &userList[j])
			}
		}
	}
}

func change(a *User, b *User) {
	temp := a.Name
	a.Name = b.Name
	b.Name = temp
}

func main() {
	list := []User{
		{"Paul", 19},
		{"John", 21},
		{"Jane", 35},
		{"Abraham", 25},
	}

	sorting(list) // list type은 pass by value or reference?
	for _, user := range list {
		fmt.Println(user.Name)
	}
}
