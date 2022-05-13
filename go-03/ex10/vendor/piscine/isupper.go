package piscine

func isUpperRune(c rune) bool {
	if	'A' <= c && c <= 'Z' {
		return true
	}
	return false
}

func IsUpper(s string) bool {
	ret := []rune(s)
	for _, r := range ret {
		if isUpperRune(r) == false {
			return false
		}
	}
	return true
}