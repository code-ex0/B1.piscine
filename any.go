package challenge_go

func Any(f func(string) bool, a []string) bool {
	for _, lettre := range a {
		if f(lettre) {
			return true
		}
	}
	return false
}
