package main

import (
	"fmt"
)

func GetnthFeb(n int) {
	a := 0
	b := 1
	fmt.Printf("The First %vth number of feb is \n", n)
	fmt.Printf("%v ", a)
	if n == 1 {
		return
	}
	fmt.Printf("%v ", b)
	for i := 2; i <= n; i++ {
		c := a + b
		if i&1 == 0 {
			a = c
		} else {
			b = c
		}
		fmt.Printf("%v ", c)
	}
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	GetnthFeb(n)
}
