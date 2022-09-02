package challenge_go

func IsSorted(f func(a, b int) int, a []int) bool {
	var ascendingordeur = true
	var descendingordeur = true
	if len(a) == 1 {
		return true
	}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			ascendingordeur = false
		}
	}
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			descendingordeur = false
		}
	}
	return ascendingordeur || descendingordeur
}
