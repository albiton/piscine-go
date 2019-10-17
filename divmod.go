package piscine

<<<<<<< HEAD
func DivMod(a int, b int, div *int, mod *int) {
	c := a / b
	d := a % b
=======
import (
     "fmt"
)

func DivMod(a int, b int, div *int, mod *int) {
	*div = a/b
	*mod = a%b
	fmt.Print(division)
	fmt.Print(modulo)
>>>>>>> bfa6d90223939b70444c1a3d664223e917ec87cf
}
