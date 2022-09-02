package challenge_go

func NRune(s string, n int) rune {
	if n <= len(s) && n > 0 {
		return []rune(s)[n-1]
	}
	return 0
}
