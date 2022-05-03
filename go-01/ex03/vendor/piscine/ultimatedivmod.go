package piscine

func UltimateDivMod(a *int, b *int) {
	var div = *a / *b
	var mod = *a % *b
	*a = div
	*b = mod
}