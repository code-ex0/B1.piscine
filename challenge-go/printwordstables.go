package challenge_go

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for j := 0; j < len(i); j++ {
			z01.PrintRune(rune(i[j]))
		}
		z01.PrintRune('\n')
	}
}
