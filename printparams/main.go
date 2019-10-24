package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args[1:] {
		for _, k := range arg {
			z01.PrintRune(k)
		}
		z01.PrintRune(10)
	}
}
