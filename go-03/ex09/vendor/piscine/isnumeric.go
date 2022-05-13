package piscine

func isNumericalRune(c rune) bool {
	if	'0' <= c && c <= '9' {
		return true
	}
	return false
}

func IsNumeric(s string) bool {
	ret := []rune(s)
	for _, r := range ret {
		if isNumericalRune(r) == false {
			return false
		}
	}
	return true
}