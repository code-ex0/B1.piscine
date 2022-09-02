package challenge_go

func IsAlpha(s string) bool {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if !(a[i] >= 'A' && a[i] <= 'Z' || a[i] >= 'a' && a[i] <= 'z' || a[i] >= '0' && a[i] <= '9') {
			return false
		}
	}
	return true
}
