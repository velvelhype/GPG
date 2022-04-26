package piscine

import "ft"

func print_digits(digits []int, n int){
	for i := 0; i < n; i++ {
		ft.PrintRune(rune('0' + digits[i]))
	}
}

func count_up(digits []int, n int, count int)  {
	digits[count]++;
	count++;
	for count < n {
		digits[count] = digits[count - 1] + 1;
		count++;
	}
}

func is_end(digits []int, n int) bool {
	for i := 0; i < n; i++ {
		if digits[i] != 10 - n + i {
			return false;
		}
	}
	return true;
}

func count_up_and_print(digits []int, n int, count int){
	for digits[count] <= (10 - n + count) {
		if (count) == n - 1{
			print_digits(digits, n);
			if is_end(digits, n) == false {
				ft.PrintRune(',');
				ft.PrintRune(' ');
			}
		} else {
			count_up_and_print(digits, n, count + 1)
		}
		count_up(digits, n, count);
	}
}

func PrintCombN(n int){
	var digits []int
	digits = make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i;
	}
	count_up_and_print(digits, n, 0)
	ft.PrintRune('\n');
}

// func PrintCombN(){
// 	var first int;
// 	var second int;
// 	var third int;

// 	first = 0;
// 	second = 1;
// 	third = 2;
// 	for (first <= 7) {
// 		for (second <= 8) {
// 			for (third <= 9){
// 				ft.PrintRune(rune('0' + first));
// 				ft.PrintRune(rune('0' + second));
// 				ft.PrintRune(rune('0' + third));
// 				if first == 7 && second == 8 && third == 9 {
// 					break;
// 				}
// 				ft.PrintRune(',')
// 				ft.PrintRune(' ')
// 				third++;
// 			}
// 			second++;
// 			third = second + 1;
// 		}
// 		first++;
// 		second = first + 1;
// 		third = second + 1;
// 	}
// 	ft.PrintRune('\n');
// }