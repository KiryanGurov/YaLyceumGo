package main

import (
	"fmt"
)

func main() {
	var sign string
	var temp int
	fmt.Scan(&sign)
	if sign == "0" {
		temp = 0
	} else {
		fmt.Scan(&temp)
		if sign == "-" {
			temp = 0 - temp
		}
	}
	switch {
	case temp > 20:
		fmt.Println("Стоит надеть майку и шорты")
	case temp > 9 && temp < 21:
		fmt.Println("Стоит надеть штаны и кофту")
	case temp > -6 && temp < 10:
		fmt.Println("Стоит надеть куртку")
	default:
		fmt.Println("Стоит надеть зимнюю куртку")
	}
}
