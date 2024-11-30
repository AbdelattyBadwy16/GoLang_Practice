package main

import (
	"fmt"
)

func Swap(a *int, b *int) {
	*a = *b + *a
	*b = *a - *b
	*a = *a - *b
}

func main() {
	var a, b int
	fmt.Printf("Enter two number : ")
	fmt.Scanf("%v %v", &a, &b)
	Swap(&a, &b)
	fmt.Printf("\n a: %v , b:%v ", a, b)
}
