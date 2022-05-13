package main

import (
"ft"
"piscine"
)

func main() {
ft.PrintRune(piscine.NRune("Hello!", 3))
ft.PrintRune(piscine.NRune("Salut!", 2))
ft.PrintRune(piscine.NRune("Bye!", -1))
ft.PrintRune(piscine.NRune("Bye!", 5))
ft.PrintRune(piscine.NRune("Ola!", 4))
ft.PrintRune(piscine.NRune("犬神家", -1))
ft.PrintRune(piscine.NRune("犬神家", 0))
ft.PrintRune(piscine.NRune("犬神家", 1))
ft.PrintRune(piscine.NRune("犬神家", 2))
ft.PrintRune(piscine.NRune("犬神家", 3))
ft.PrintRune(piscine.NRune("犬神家", 4))
ft.PrintRune('\n')
}