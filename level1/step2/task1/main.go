package main

import "fmt"

func main() {

    var x, y, z float64
    fmt.Scan(&x, &y, &z)
    if x == y && y == z && z == x {
        fmt.Println("Максимальное равенство")
        return
    }
    fmt.Println("Не равны")
}