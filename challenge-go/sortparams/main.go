package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args
	length := len(args)
	for i := 1; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}

	for i := 1; i < length; i++ {
		for _, letter := range args[i] {
			z01.PrintRune(letter)
		}
		z01.PrintRune('\n')
	}
}
