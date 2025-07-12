package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scan(&number)

	if sqrt := math.Sqrt(number); !math.IsNaN(sqrt) {
		fmt.Printf("%.3f", sqrt)
		return
	}
	fmt.Println(-1)
}
