package challenge_go

func ToLower(s string) string {
	var temps string
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] >= 'A' && a[i] <= 'Z' {
			temps += string(a[i] + 32)
		} else {
			temps += string(a[i])
		}
	}
	return temps
}
