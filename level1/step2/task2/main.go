package main
import "fmt"

func main() {
    var pwd1, pwd2 string
    fmt.Scan(&pwd1, &pwd2)

    len1, len2 := len(pwd1), len(pwd2)
    switch{
    case len1 > 7 && len2 > 7:
        fmt.Println("Оба пароля надёжные")
    case len1 < 8 && len2 < 8:
        fmt.Println("Оба пароля ненадёжные")
    case len1 > 7:
        fmt.Println("Только первый пароль надёжный")
    case len2 > 7:
        fmt.Println("Только второй пароль надёжный")
    }
}