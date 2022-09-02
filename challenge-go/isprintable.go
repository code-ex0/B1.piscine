package challenge_go

func IsPrintable(s string) bool {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if !(a[i] >= ' ' && a[i] <= '~') {
			return false
		}
	}
	return true
}
