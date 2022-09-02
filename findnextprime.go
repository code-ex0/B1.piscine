package challenge_go

func FindNextPrime(nb int) int {
	if nb > 0 {
		if nb <= 1 {
			return 2
		}
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				return FindNextPrime(nb + 1)
			}
		}
		return nb
	} else {
		return 2
	}
}
