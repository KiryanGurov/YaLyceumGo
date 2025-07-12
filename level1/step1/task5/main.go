package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	switch {
	case x == 0:
		fmt.Println("Число 0")
	case -10 < x && x < 10:
		fmt.Println("Число однозначное")
	case x % 2 == 0:
		fmt.Println("Число чётное")
	case x > 0:
		fmt.Println("Число положительное")
	default:
		fmt.Println("Число красивое")
	}
}
