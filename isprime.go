package challenge_go

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb <= 1 {
			return false
		}
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}
