package main
import (
"fmt"
"piscine"
)
func main() {
fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
fmt.Println(piscine.Capitalize("hello! How are you? How+are+things+4you?"))
fmt.Println(piscine.Capitalize("   Hello! How are you? How+are+things+4you?"))
fmt.Println(piscine.Capitalize("GEKKO! How are you? How+are+things+4you?"))
}