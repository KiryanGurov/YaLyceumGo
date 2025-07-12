package main

import "fmt"

func is_winner(pl1, pl2 string) bool {
	if (pl1 == "камень" && pl2 == "ножницы") || (pl1 == "ножницы" && pl2 == "бумага") || (pl1 == "бумага" && pl2 == "камень") {
		return true
	}
	return false
}

func main() {
	var ch1, ch2 string
	fmt.Scan(&ch1, &ch2)

	if is_winner(ch1, ch2) {
		fmt.Println("Первый игрок победил")
	} else if is_winner(ch2, ch1) {
		fmt.Println("Второй игрок победил")
	} else {
		fmt.Println("Ничья")
    }
      
}
