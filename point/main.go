package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbr(n int) {
	TempInt := n
	var resulta []int
	for TempInt > 0 {
		resulta = append(resulta, TempInt%10)
		TempInt /= 10
	}
	if n > 0 {
		for i := 0; i < len(resulta); i++ {
			z01.PrintRune(rune(resulta[len(resulta)-i-1]) + 48)
		}
	} else {
		z01.PrintRune('0')
	}

}

func main() {
	text1 := "x = "
	text2 := ", y = "
	points := &point{}

	setPoint(points)

	for _, l := range text1 {
		z01.PrintRune(l)
	}
	PrintNbr(points.x)
	for _, l := range text2 {
		z01.PrintRune(l)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
