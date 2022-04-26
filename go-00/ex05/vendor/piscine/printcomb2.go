package piscine

import "ft"


func act2(act1_first rune, act1_second rune, first int, second int) {
	for (first <= 9) {
		for (second <= 9) {
			ft.PrintRune(act1_first);
			ft.PrintRune(act1_second);
			ft.PrintRune(' ')
			ft.PrintRune(rune('0' + first));
			ft.PrintRune(rune('0' + second));
			if first == 9 && second == 9 {
				break;
			}
			ft.PrintRune(',')
			ft.PrintRune(' ')
			second++;
		}
		first++;
		second = 0;
	}
}

func PrintComb2(){
	var first int;
	var second int;

	first = 0;
	second = 0;
	for first <= 9{
		for second <= 9{
				act2(rune('0' + first), rune('0' + second), first, second + 1);
				if first == 9 && second == 8 {
					break;
				}
				ft.PrintRune(',')
				ft.PrintRune(' ')
			second++;
		}
		first++;
		second = 0;
	}
	ft.PrintRune('\n');
}