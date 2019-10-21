package piscine

func BasicAtoi2(s string) int {
	chislo1 := 0
	for _, cifra1 := range s {
		if cifra1 < '0' || cifra1 > '9' {
			return 0
		}
		for i := 0; i <= 9; i += 1 {
			if rune(i) == cifra1-'0' {
				chislo1 = chislo1*10 + i
				break
			}
		}
	}
	return chislo1
}
