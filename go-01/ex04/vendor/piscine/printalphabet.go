package piscine

import "ft"

func PrintStr(s string) {
	for _, i := range s {
		ft.PrintRune(i)
	}
}