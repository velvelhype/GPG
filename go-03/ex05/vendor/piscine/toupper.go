package piscine

func ToUpper(s string) string {
	S := []rune(s)
	for i, r := range S {
		if 'a' <= r && r <= 'z' {
			S[i] += 'A' - 'a'
		}
	}
	return  string(S)
}