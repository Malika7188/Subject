package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	result := []string{}
	start := 0
	for i := 0; i < len(args); i++ {
		if args[i] == ' ' {
			result = append(result, args[start:i])
			start = i + 1
		}
		if i == len(args)-1 {
			result = append(result, args[start:])
		}
	}
	for i := len(result) - 1; i >= 0; i-- {
		if i == len(result) {
			fmt.Print(result[i])
		} else {
			fmt.Print(result[i] + " ")
		}
	}
	fmt.Println()
}
