package challenge_go

func ListSize(l *List) (i int) {
	it := l.Head
	for it != nil {
		i++
		it = it.Next
	}
	return
}
