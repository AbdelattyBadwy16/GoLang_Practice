package main

import (
	"fmt"
	"math"
)

func Order(n int) int {
	cnt := 0
	for n > 0 {
		cnt++
		n /= 10
	}
	return cnt
}

func IsArmStrong(n int) bool {
	sum := 0
	ord := Order(n)
	cop := n
	for n > 0 {
		dig := n % 10
		sum += int(math.Pow(float64(dig), float64(ord)))
		n /= 10
	}
	return sum == cop
}

func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%v", &n)
	if IsArmStrong(n) == true {
		fmt.Println("The Nums Is ArmStrong.")
	} else {
		fmt.Println("The Nums Is Not ArmStrong.")
	}
}
