package main

import piscine "github.com/01-edu/z01"

func PrintComb() {

	a := 1
	b := 2
	for i := 0; i < =7; i++ {
		for a = i + 1; a <= 8; j++ {
			for b = a + 1 ; b <= 9; b++ {
				z01.Print(i)
				z01.Print(a)
				z01.Print(b)
				if i < 7 {
					z01.Print(",")
				} else {
					z01.Print(10)
				} 
			}
		}
		
	}
}