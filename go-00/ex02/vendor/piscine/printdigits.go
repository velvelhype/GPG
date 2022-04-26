package piscine

import "ft"

func PrintDigits() {
	c := '0';
	for
	{
		ft.PrintRune(c);
		if c == '9' {
			break
		}
		c++;
	}
	ft.PrintRune('\n');
}