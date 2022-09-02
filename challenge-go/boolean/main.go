package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if !(nbr%2 == 1) {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := os.Args[1:]
	if isEven(len(lengthOfArg)) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
