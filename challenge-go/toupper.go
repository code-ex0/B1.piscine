package challenge_go

func ToUpper(s string) string {
	var temps string
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if a[i] >= 'a' && a[i] <= 'z' {
			temps += string(a[i] - 32)
		} else {
			temps += string(a[i])
		}
	}
	return temps
}
