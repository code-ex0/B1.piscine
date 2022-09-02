package challenge_go

func IterativePower(nb int, power int) int {
	var a int = 1
	if power >= 0 {
		for i := 1; i <= power; i++ {
			a *= nb
		}
	} else {
		return 0
	}
	return a
}
