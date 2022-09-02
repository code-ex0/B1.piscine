package challenge_go

func ListMerge(l1 *List, l2 *List) {

	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	} else if l2.Head == nil {
		return
	}
	it2 := l2.Head
	it1 := l1.Tail
	for it2 != nil {
		it1.Next = it2
		it2 = it2.Next
		it1 = it1.Next
	}
}
