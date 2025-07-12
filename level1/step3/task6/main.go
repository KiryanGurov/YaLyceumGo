package main

import (
	"fmt"
)

func main() {
	var (
		inputWord  string
		resultWord string
	)
	fmt.Scanln(&inputWord)

	for _, letter := range inputWord {
		if letter == 'а' || letter == 'о' {
			continue
		}
		resultWord += string(letter)
	}

	fmt.Println(resultWord)
}
