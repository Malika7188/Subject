package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	str1 := args[0]
	str2 := args[1]

	if wdmatch(str1, str2) {
		for _, char := range str1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}

func wdmatch(str1, str2 string) bool {
	var i, j int
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}
	if i == len(str1) {
		return true
	} else {
		return false
	}
}
