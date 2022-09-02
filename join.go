package challenge_go

func Join(strs []string, sep string) string {
	var a string
	for i := 0; i < len(strs); i++ {
		if i != len(strs) && i != 0 {
			a += sep + strs[i]
		} else {
			a += strs[i]
		}
	}
	return a
}
