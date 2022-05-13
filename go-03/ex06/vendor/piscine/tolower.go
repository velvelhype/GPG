package piscine

func ToLower(s string) string {
	S := []rune(s)
	for i, r := range S {
		if 'A' <= r && r <= 'Z' {
			S[i] += 'a' - 'A'
		}
	}
	return  string(S)
}