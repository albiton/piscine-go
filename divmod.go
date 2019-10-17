package piscine

import "fmt"

func DivMod(a int, b int, division *int, modulo *int) {
	*division = a / b
	*modulo = a % b
	fmt.Print(division)
	fmt.Print(modulo)
}
