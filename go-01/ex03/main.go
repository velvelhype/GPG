package main

import (
"fmt"
"piscine"
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(nil, &b)
	fmt.Println(a)
	fmt.Println(b)
}
