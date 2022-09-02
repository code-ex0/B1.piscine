package main

import (
	"github.com/01-edu/z01"
	"os"
)

func BasicAtoi(s string) int {
	runes := []rune(s)
	var Numbre int = 0
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 48 && runes[i] <= 57 {
			Numbre *= 10
			Numbre += int(runes[i]) - 48
		}
	}
	return Numbre
}

func main() {
	args := os.Args
	length := len(args)
	add := 96
	start := 1
	letter := 0

	if length > 1 {
		if args[1] == "--upper" {
			add = 64
			start = 2
		}
		for i := start; i < length; i++ {
			letter = BasicAtoi(args[i])
			if letter > 0 && letter <= 26 {
				z01.PrintRune(rune(BasicAtoi(args[i]) + add))
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
