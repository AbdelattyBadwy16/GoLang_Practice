package main

import (
	"fmt"
)

func IsPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	if IsPrime(n) == true {
		fmt.Printf("%v is a Prime number: ", n)
	} else {
		fmt.Printf("%v is not a Prime number: ", n)
	}
}
