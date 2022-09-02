package challenge_go

func Atoi(s string) int {
	runes := []rune(s)
	var Numbre int = 0
	Sigle := 2
	IsNumber := true
	Mult := 1
	HaveNumber := false
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 48 && runes[i] <= 57 {
			Numbre *= 10
			Numbre += int(runes[i]) - 48
			HaveNumber = true
		} else if runes[i] == 43 || runes[i] == 45 {
			Sigle--
			HaveNumber = false
			if runes[i] == 45 {
				Mult = -1
			}
		} else {
			IsNumber = false
		}
	}
	if Sigle > 0 && IsNumber == true && HaveNumber == true {
		return Numbre * Mult
	} else {
		return 0
	}
}
