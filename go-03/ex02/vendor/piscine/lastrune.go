package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func LastRune(s string) rune {
	n := StrLen(s)
	return ([]rune(s)[n - 1])
}