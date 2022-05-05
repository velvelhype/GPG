package piscine

import (
	"ft"
)

func EightQueens() {
	var array [8]int
	rec(0, array)
}

func rec(i int, array [8]int) {
	if i == 8 {
		printArray(array)
		return 
	}
	for j := 0; j < 8; j++ {
		if is_valid(i, array, j) == true {
			array[i] = j
			rec(i + 1, array)
		}
	}
}

func printArray(array [8]int) {
	for i := 0; i < 8; i++{
		ft.PrintRune(rune('0' + array[i] + 1))
	}
	ft.PrintRune(rune('\n'))
}

func is_valid(count int, array [8]int, cur int)bool {
	if	count == 0 {
		return true
	}
	for i := count - 1; i >= 0; i-- {
		// is checkable wide
		if array[i] == cur {
			return false
		}
		// is checkable tiltly
		if array[i]  + (count - i) == cur || array[i] - + (count - i) == cur {
			return false
		}
	} 
	return true
}