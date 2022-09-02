package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for j := len(argsWithoutProg) - 1; j >= 0; j-- {
		for i := 0; i < len(argsWithoutProg[j]); i++ {
			z01.PrintRune(rune(argsWithoutProg[j][i]))
		}
		z01.PrintRune('\n')
	}
}
