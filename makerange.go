package challenge_go

func MakeRange(min, max int) []int {
	var temp []int
	lens := max - min
	if max-min > 0 {
		temp = make([]int, lens)
		for i := 0; i < lens; i++ {
			temp[i] = i + min
		}
	}
	return temp
}
