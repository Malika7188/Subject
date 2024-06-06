package main

import (
	"revision"

	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(revision.LastRune("Hello!"))
	z01.PrintRune(revision.LastRune("Salut!"))
	z01.PrintRune(revision.LastRune("Ola!"))
	z01.PrintRune('\n')
}
