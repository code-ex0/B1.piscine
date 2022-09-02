package challenge_go

func AppendRange(min, max int) []int {
	var temp []int
	if max-min >= 0 {
		for i := min; i < max; i++ {
			temp = append(temp, i)
		}
	}
	return temp
}
