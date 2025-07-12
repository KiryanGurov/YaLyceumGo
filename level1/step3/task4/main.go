package main

import (
	"fmt"
)

func main() {
	var answer string
	for {
		fmt.Scanln(&answer)
		switch answer {
		case "да", "нет", "чёрный", "белый":
			fmt.Println("Поражение")
			return
		default:
			fmt.Println("Игра продолжается")
		}
	}
}
