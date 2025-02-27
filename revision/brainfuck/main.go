package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]
	memery := make([]byte, 2048)
	bracket := make(map[int]int)
	pointer := 0
	stack := []int{}

	for i, c := range arg {
		if c == '[' {
			stack = append(stack, i)
		}
		if c == ']' {
			opening := stack[len(stack)-1]
			bracket[opening] = i
			bracket[i] = opening
			stack = stack[:len(stack)-1]
		}

	}
	for i := 0; i < len(arg); i++ {
		if arg[i] == '>' {
			pointer++
		}
		if arg[i] == '<' {
			pointer--
		}
		if arg[i] == '+' {
			memery[pointer]++
		}
		if arg[i] == '-' {
			memery[pointer]--
		}
		if arg[i] == '.' {
			z01.PrintRune(rune(memery[pointer]))
		}
		if arg[i] == '[' && memery[pointer] == 0 {
			i = bracket[i]
		}
		if arg[i] == ']' && memery[pointer] != 0 {
			i = bracket[i]
		}
	}

}
