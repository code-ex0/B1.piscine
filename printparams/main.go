package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	for _, ParamPrint := range argsWithoutProg {
		for j := 0; j < len(ParamPrint); j++ {
			z01.PrintRune(rune(ParamPrint[j]))
		}
		z01.PrintRune('\n')
	}
}
