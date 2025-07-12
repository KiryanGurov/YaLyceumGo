package main

import (
	"fmt"
)

func main() {
	var roubl, kop, price int
	fmt.Scan(&roubl, &kop, &price)

	if total := roubl + kop/100.0; total >= price {
		fmt.Println("Сегодня будет вкусный кофе!")
		return
	}
	fmt.Println("Стоит подкопить")
}
