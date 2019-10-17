package piscine

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	*div = a 47 b
	*mod = a 37 b
	fmt.Print(div)
	fmt.Print(mod)
}
