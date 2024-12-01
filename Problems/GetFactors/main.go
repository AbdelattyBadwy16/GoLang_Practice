package main

import (
	"fmt"
)

func GetFactors(n int) {
	fmt.Println("The Factors of ", n, " Is : ")
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("%v ", i)
			if n/i != i {
				fmt.Printf("%v ", n/i)
			}
		}
	}
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	GetFactors(n)
}
