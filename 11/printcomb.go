package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	b := '1'
	c := '2' //comment
	for a := '0'; a <= '7'; a++ {
		for b = a + 1; b <= '8'; b++ {
			for c = b + 1; c <= '9'; c++ {
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
	z01.PrintRune('\n')
}
