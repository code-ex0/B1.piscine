package main

import (
	"github.com/01-edu/z01"
	"io/ioutil"
	"os"
)

func PrintWitheZ01(s string) {
	for _, Lettre := range s {
		z01.PrintRune(Lettre)
	}
}

func PrintFatal(s string) {
	PrintWitheZ01(s + "\n")
	os.Exit(1)
}

func main() {
	i := len(os.Args)
	if i == 1 {
		temp, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			PrintFatal(err.Error())
		}
		PrintWitheZ01(string(temp))
	} else {
		for _, file := range os.Args[1:] {
			temp, err := ioutil.ReadFile(file)
			if err != nil {
				PrintFatal("ERROR: " + err.Error())
			}
			PrintWitheZ01(string(temp))
		}
	}
}
