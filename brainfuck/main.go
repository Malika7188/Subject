package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	memory := make([]byte, 2048)
	brackets := make(map[int]int)
	pointer := 0 
	stack := []int{}

	for i, c := range args {
		if c == '[' {
			stack = append(stack, i)
		}
		if c == ']' {
			opening := stack[len(stack)-1]
			brackets[opening] = i
			brackets[i] = opening
			stack = stack[:len(stack)-1]
		}
	}
	for i := 0; i < len(args); i++ {
		if args[i] == '>' {
			pointer++
		}
		if args[i] == '<' {
			pointer--
		}
		if args[i] == '+' {
			memory[pointer]++
		}
		if args[i] == '-' {
			memory[pointer]--
		}
		if args[i] == '.' {
			z01.PrintRune(rune(memory[pointer]))
		}
		if args[i] == '[' && memory[pointer] == 0 {
			i = brackets[i]
		}
		if args[i] == ']' && memory[pointer] != 0 {
			i = brackets[i]
		}
	}
}