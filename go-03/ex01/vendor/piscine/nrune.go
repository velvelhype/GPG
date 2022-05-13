package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func NRune(s string, n int) rune {
	if n <= 0 || n > StrLen(s) {
		return (0)
	}
	return ([]rune(s)[n - 1])
}