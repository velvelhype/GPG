package piscine

func isAlpha(c rune) bool {
	if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z'){
		return true
	}
	return false
}

func Capitalize(s string) string {
	CS := []rune(s)
	ini_word_flag := true
	for i, r := range CS {
		switch {
		case isAlpha(r) == false:
			ini_word_flag = true
		case ini_word_flag == true:
			if 'a' <= r && r <= 'z'{
				CS[i] += 'A' - 'a'
			}
			ini_word_flag = false
		case ini_word_flag == false && 'A' <= r && r <= 'Z':
			CS[i] += 'a' - 'A'
		}
	}
	return  string(CS)
}