package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if s == "Go" {
		fmt.Println("Go!")
		return
	}
	fmt.Println("Я знаю только Go!")
}
