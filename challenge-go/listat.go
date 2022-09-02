package challenge_go

func ListAt(l *NodeL, pos int) *NodeL {
	it := l
	var temp = &NodeL{}

	if pos < LenNodeL(it) {
		for i := 0; i < pos+1; i++ {
			temp = it
			it = it.Next
		}
	} else {
		temp = nil
	}
	return temp
}

func LenNodeL(l *NodeL) (Result int) {
	for l != nil {
		Result++
		l = l.Next
	}
	return
}
