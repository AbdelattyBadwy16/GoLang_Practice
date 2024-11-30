package main

import (
	"fmt"
)

func LargeNum(a int, b int, c int) int {
	if a > b && a > c {
		return a
	}
	if b > a && b > c {
		return b
	}
	return c
}

func main() {
	var a, b, c int
	fmt.Printf("Enter the 3 numbers : ")
	fmt.Scanf("%v %v %v", &a, &b, &c)
	fmt.Printf("The Large Num is : %v", LargeNum(a, b, c))
}
