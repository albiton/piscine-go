package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	ProgrammName := os.Args

	for _, i := range ProgrammName[0] {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
