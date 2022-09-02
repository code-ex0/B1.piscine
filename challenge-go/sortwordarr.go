package challenge_go

func SortWordArr(a []string) {
	args := make([]string, len(a))
	copy(args, a)
	length := len(args)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	copy(a, args)
}
