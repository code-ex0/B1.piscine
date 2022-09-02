package challenge_go

func ConcatParams(args []string) string {
	var temp string
	for i := 0; i < len(args); i++ {
		if !(i == len(args)-1) {
			temp += args[i] + "\n"
		} else {
			temp += args[i]
		}
	}
	return temp
}
