package challenge_go

func BasicAtoi(s string) int {
	runes := []rune(s)
	var Numbre int = 0
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 48 && runes[i] <= 57 {
			Numbre *= 10
			Numbre += int(runes[i]) - 48
		}
	}
	return Numbre
}
