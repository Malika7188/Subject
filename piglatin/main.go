package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	result := ""

	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for i := 0; i < len(arg); i++ {
		for _, vow := range vowels {
			if arg[0] == vow {
				result = arg + "ay"
				break
			}
			if arg[i] == vow {
				result = arg[i:] + arg[:i] + "ay"
				break
			}
		}
	}
	if result == "" {
		fmt.Println("No vowels")
	} else {
		fmt.Println(result)
	}
}
