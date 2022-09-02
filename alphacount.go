package challenge_go

func AlphaCount(s string) int {
	a := []rune(s)
	temps := 0
	for i := 0; i < len(a); i++ {
		if a[i] >= 'A' && a[i] <= 'Z' || a[i] >= 'a' && a[i] <= 'z' {
			temps++
		}
	}
	return temps
}
