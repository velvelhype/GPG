package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

func IndexByte(s string, c byte) int {
	for i := 0; i <= StrLen(s) - 1; i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

func Index(s string, toFind string) int {
	n := StrLen(toFind)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexByte(s, toFind[0])
	case n == StrLen(s):
		if toFind == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	c0 := toFind[0]
	i := 0
	t := StrLen(s) - n + 1
	for i < t {
		if s[i] != c0 {
			o := IndexByte(s[i+1:t], c0)
			if o < 0 {
				return -1
			}
			i += o + 1
		}
		if s[i:i+n] == toFind {
			return i
		}
		i++
	}
	return -1
}