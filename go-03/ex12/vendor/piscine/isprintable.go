package piscine

func isPrintableRune(c rune) bool {
	if	' ' <= c && c <= '~' {
		return true
	}
	return false
}

func IsPrintable(s string) bool {
	ret := []rune(s)
	for _, r := range ret {
		if isPrintableRune(r) == false {
			return false
		}
	}
	return true
}