package main

import (
	"fmt"
)

func sumArray(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func main() {
	var count, discount int
	fmt.Scan(&count, &discount)
	var costs = make([]float64, count)
	for i := range count {
		fmt.Scan(&costs[i])
	}

	dc := 1 - float64(discount)/100.0
	var sum float64
	for _, cost := range costs {
		sum += cost * dc
	}
	fmt.Println(sum)
}
