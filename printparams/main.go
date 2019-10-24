package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, res := range args[1:] {
		for _, k := range res {
			z01.PrintRune(k)
		}
		z01.PrintRune(10)
	}
}
