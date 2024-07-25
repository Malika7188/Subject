package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	val1, err := strconv.Atoi(os.Args[1])
	checkerr(err)
	val2, err := strconv.Atoi(os.Args[2])
	checkerr(err)

	for i := val1; i <= val2; i++ {
		PrintS(strconv.Itoa(i))
	}
	for i := val1; i >= val2; i-- {
		PrintS(strconv.Itoa(i))
	}
	z01.PrintRune('\n')
}

func PrintS(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}

func checkerr(err error) {
	if err != nil {
		return
	}
}
