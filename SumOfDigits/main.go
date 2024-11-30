package main

import (
	"fmt"
)

func SumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	fmt.Printf("The Sum of Digits is : %v", SumOfDigits(n))
}
