package challenge_go

func BasicJoin(elems []string) string {
	var a string
	for i := 0; i < len(elems); i++ {
		a += elems[i]
	}
	return a
}
