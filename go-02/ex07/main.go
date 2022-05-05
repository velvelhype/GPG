package main
import (
        "fmt"
        "piscine"
)
func main() {
        fmt.Println(piscine.FindNextPrime(1000000004))
        fmt.Println(piscine.FindNextPrime(12))
        fmt.Println(piscine.FindNextPrime(5))
        fmt.Println(piscine.FindNextPrime(4))
        fmt.Println(piscine.FindNextPrime(-1))
        fmt.Println(piscine.FindNextPrime(-100))
}