package challenge_go

func Swap(a *int, b *int) {
	tempsA := *a
	tempsB := *b
	*b = tempsA
	*a = tempsB
}
