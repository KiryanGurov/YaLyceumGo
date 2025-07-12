package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scanln(&count)
	var nums = make([]float64, count)
	for i := range count {
		fmt.Scanln(&nums[i])
	}

	for i := range count {
		switch {
		case 90.0 <= nums[i] && nums[i] <= 100.0:
			fmt.Println(5)
		case 75.0 <= nums[i] && nums[i] <= 89.0:
			fmt.Println(4)
		case 50.0 <= nums[i] && nums[i] <= 74.0:
			fmt.Println(3)
		case 0.0 <= nums[i] && nums[i] <= 49.0:
			fmt.Println(2)
		default:
			fmt.Println("Неверный балл")
		}
	}
}
