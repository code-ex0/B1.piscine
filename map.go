package challenge_go

func Map(f func(int) bool, a []int) []bool {
	results := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		results[i] = f(a[i])
	}
	return results
}
