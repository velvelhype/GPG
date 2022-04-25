package piscine

import "ft"

func PrintAlphabet() {
	c := 'a';
	for
	{
		ft.PrintRune(c);
		if c == 'z' {
			break
		}
		c++;
	}
}