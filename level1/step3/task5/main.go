package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		totalWords int
		searchWord string
	)
	fmt.Scan(&totalWords, &searchWord)

	var (
		word       string
		countWords int
	)
	searchWord = strings.ToLower(searchWord)
	for range totalWords {
		fmt.Scan(&word)
		if searchWord == strings.ToLower(word) {
			countWords++
		}
	}

	fmt.Println(countWords)
}
