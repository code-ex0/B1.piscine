package challenge_go

func IterativeFactorial(nb int) int {
	result := 1
	if nb <= 50 && nb >= 0 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
	} else {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if result < 0 {
		return 0
	}
	return result
}
