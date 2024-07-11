package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1]

	if len(os.Args) != 2 {
		return
	}

	check := false
	for _, char := range args {
		if char != ' ' {
			check = true
			break
		}
	}
	if check {

		word := ""
		for i := len(args) - 1; i >= 0; i-- {
			if word == "" && args[i] == ' ' {
				continue
			}
			if word != "" && args[i] == ' ' {
				break
			}
			word = string(args[i]) + word
		}
		fmt.Println(word)
	}
}
