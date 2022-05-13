package main

import (
"fmt"
"piscine"
"strings"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(strings.Index("Hello!", "l"))

	fmt.Println(piscine.Index("Hello!", "Hello!!"))
	fmt.Println(strings.Index("Hello!", "Hello!!"))

	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(strings.Index("Salut!", "alu"))

	fmt.Println(piscine.Index("Ola!", "hOl"))
	fmt.Println(strings.Index("Ola!", "hOl"))

	fmt.Println(piscine.Index("ABC", "A"))
	fmt.Println(strings.Index("ABC", "A"))

	fmt.Println(piscine.Index("ABC", "B"))
	fmt.Println(strings.Index("ABC", "B"))

	fmt.Println(piscine.Index("ABC", "C"))
	fmt.Println(strings.Index("ABC", "C"))

	fmt.Println(piscine.Index("ABC", "D"))
	fmt.Println(strings.Index("ABC", "D"))

	fmt.Println(piscine.Index("ABC", ""))
	fmt.Println(strings.Index("ABC", ""))

	fmt.Println(piscine.Index("", ""))
	fmt.Println(strings.Index("", ""))
}