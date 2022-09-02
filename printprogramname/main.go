package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	var Index int
	s := os.Args[0]
	for i := 0; i < len(s); i++ {
		if s[i] == 92 {
			Index = i
		}
	}
	for _, LettrePrint := range s[Index+1:] {
		z01.PrintRune(LettrePrint)
	}
	z01.PrintRune('\n')
}
