package piscine

func isAlphaNumericalRune(c rune) bool {
	if	('a' <= c && c <= 'z')	||
		('A' <= c && c <= 'Z' )	|| 
		('0' <= c && c <= '9' ){
		return true
	}
	return false
}

func IsAlpha(s string) bool {
	ret := []rune(s)
	for _, r := range ret {
		if isAlphaNumericalRune(r) == false {
			return false
		}
	}
	return true
}