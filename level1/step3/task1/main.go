package main

import (
	"fmt"
)

func main() {
	var s [5]string
	for i := range 5 {
		fmt.Scanln(&s[i])
	}
	for i := range 5 {
		if s[i] == "Go" {
			fmt.Println("Go!")
		} else {
			fmt.Println("Я знаю только Go!")
		}
	}
}
