package challenge_go

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	Current, TempList := l.Head, l.Head
	Next, Previous := &NodeL{}, &NodeL{}

	Next = Current.Next
	Current.Next = Previous
	Previous = Current
	Current = Next

	Next = Current.Next
	Current.Next = Previous
	Previous = Current
	Current = Next

	Next = Current.Next
	Current.Next = Previous
	Previous = Current
	Current = Next

	Next = Current.Next
	Current.Next = Previous
	Previous = Current
	Current = Next

	l.Head = Previous
	l.Tail = TempList

}

/*
func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	Current, TempList := l.Head, l.Head
	Next, Previous := &NodeL{}, &NodeL{}

	for Current != nil {
		Next = Current.Next
		Current.Next = Previous
		Previous = Current
		Current = Next
	}
	l.Head = Previous
	l.Tail = TempList
	l.Tail.Next = nil
}
*/
