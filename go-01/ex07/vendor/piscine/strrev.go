package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func StrRev(s string) string {
	var rev_s string
	for _, c := range s {
		rev_s = string(c) + rev_s
	}
	return rev_s
}