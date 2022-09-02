package challenge_go

func IsNumeric(s string) bool {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if !(a[i] >= '0' && a[i] <= '9') {
			return false
		}
	}
	return true
}
