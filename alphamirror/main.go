package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]

	var result string

	for _, char := range args {
		if char >= 'a' && char <= 'z' {
			char = 'z' - char + 'a'
			result += string(char)
		} else if char >= 'A' && char <= 'Z' {
			char = 'Z' - char + 'A'
			result += string(char)
		} else {
			result += string(char)
		}
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'))  {
			fmt.Println(char)
		}
	}
	fmt.Println(result)
}
