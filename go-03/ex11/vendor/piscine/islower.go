package piscine

func isLowerRune(c rune) bool {
	if	'a' <= c && c <= 'z' {
		return true
	}
	return false
}

func IsLower(s string) bool {
	ret := []rune(s)
	for _, r := range ret {
		if isLowerRune(r) == false {
			return false
		}
	}
	return true
}