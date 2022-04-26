package piscine

import "ft"

func PrintComb(){
	var first int;
	var second int;
	var third int;

	first = 0;
	second = 1;
	third = 2;
	for (first <= 7) {
		for (second <= 8) {
			for (third <= 9){
				ft.PrintRune(rune('0' + first));
				ft.PrintRune(rune('0' + second));
				ft.PrintRune(rune('0' + third));
				if first == 7 && second == 8 && third == 9 {
					break;
				}
				ft.PrintRune(',')
				ft.PrintRune(' ')
				third++;
			}
			second++;
			third = second + 1;
		}
		first++;
		second = first + 1;
		third = second + 1;
	}
	ft.PrintRune('\n');
}