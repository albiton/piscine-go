package piscine

import (
     "fmt"
)

func DivMod(a int, b int, div *int, mod *int) {
	*div = a/b
	*mod = a%b
	fmt.Print(division)
	fmt.Print(modulo)
}
