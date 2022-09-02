package challenge_go

func ListLast(l *List) interface{} {
	it := l.Head
	var temp interface{}
	for it != nil {
		temp = it.Data
		it = it.Next
	}
	return temp
}
