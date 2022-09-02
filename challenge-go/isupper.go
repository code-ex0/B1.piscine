package challenge_go

func IsUpper(s string) bool {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if !(a[i] >= 'A' && a[i] <= 'Z') {
			return false
		}
	}
	return true
}
