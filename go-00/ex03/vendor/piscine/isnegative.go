package piscine

import "ft"

func IsNegative(nb int){
	if nb < 0 {
		ft.PrintRune('T')
		ft.PrintRune('\n')
	} else {
		ft.PrintRune('F')
		ft.PrintRune('\n')
	}
}