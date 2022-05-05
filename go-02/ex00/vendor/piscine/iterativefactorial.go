package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 || nb > 23{
		return 0
	}
	for i := nb - 1; i > 1; i--{
		nb = nb * i
	} 
	return nb
}