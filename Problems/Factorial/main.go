package main

import (
	"fmt"
)

func Fact(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	fmt.Printf("The Fact of %v is : %v", n, Fact(n))
}
