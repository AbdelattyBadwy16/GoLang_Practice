package main

import (
	"fmt"
)

func Reverse(n int) int {
	newN := 0
	for n > 0 {
		newN = (newN * 10) + n%10
		n /= 10
	}
	return newN
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	revN := Reverse(n)
	if n == revN {
		fmt.Println("The Num Is Plain.")
	} else {
		fmt.Println("The Num Is Not Plain.")
	}
}
