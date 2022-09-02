package challenge_go

func ListRemoveIf(l *List, data_ref interface{}) {

	var temp *NodeL = nil
	it := l.Head
	for it != nil {
		if !CompStr(data_ref, it.Data) {
			temp = it
			it = it.Next
			continue
		}
		it = it.Next
		if temp == nil {
			l.Head = it
		} else {
			temp.Next = it
		}
	}
}
