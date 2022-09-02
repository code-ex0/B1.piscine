package challenge_go

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var resulta []int
	if n <= 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		resulta = append(resulta, n%10)
		n /= 10
	}
	SortIntegerTable(resulta)
	for i := 0; i < len(resulta); i++ {
		z01.PrintRune(rune(resulta[i] + 48))
	}
}
