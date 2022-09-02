package challenge_go

func StrRev(s string) string {
	t := ""
	for i := len(s); i > 0; i-- {
		t += string(rune(s[i-1]))
	}
	s = t
	return s
}
