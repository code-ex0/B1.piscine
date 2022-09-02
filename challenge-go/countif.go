package challenge_go

func CountIf(f func(string) bool, tab []string) int {
	number := 0
	for _, lettre := range tab {
		if f(lettre) {
			number++
		}
	}
	return number
}
