package challenge_go

func IsLower(s string) bool {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if !(a[i] >= 'a' && a[i] <= 'z') {
			return false
		}
	}
	return true
}
