package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	if nb == 0  && nb < 20 {
		return 1
	}
	return nb * IterativeFactorial(nb-1)
	}
}
