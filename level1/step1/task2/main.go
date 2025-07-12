package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	switch {
	case x > y:
		fmt.Println("Первое число больше второго")
	case x < y:
		fmt.Println("Второе число больше первого")
	default:
		fmt.Println("Числа равны")
	}
}
