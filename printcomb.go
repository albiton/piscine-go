package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	b := 1
	c := 2
	for a := 0; a <= 7; a++ {
		for b = b + 1; b <= 8; b++ {
			for c = c + 1; c <= 9; c++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
				if a < 55 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
}
