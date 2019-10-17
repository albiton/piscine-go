package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	j := 1
	k := 2
	for i := 0; i <= 7; i++ {
		for j = i + 1; j <= 8; j++ {
			for k = j + 1; k <= 9; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i < 7 {
					z01.PrintRune(", ")
				} else {
					z01.PrintRune(10)			
						      }
			}
		}
	}
}
