package main

import (
	"fmt"
)

func IsOdd(n int) bool {
	return n&1 == 1	
}

func main(){
	var n int
	fmt.Printf("Enter The Number : ")
	fmt.Scanf("%v",&n)
	fmt.Println(n)
	if IsOdd(n) == true {
		fmt.Printf("The number %v is Odd.",n)
	}else {
		fmt.Printf("The number %v is Even.",n)
	}
}