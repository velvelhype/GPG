package piscine

import "ft"

func PrintReverseAlphabet() {
	c := 'z';
	for
	{
		ft.PrintRune(c);
		if c == 'a' {
			break
		}
		c--;
	}
	ft.PrintRune('\n');
}