package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		return
	}
	args := os.Args[1:]

	for _, str := range args {
		new := ""
		for _, c := range str {
			if c == '{' || c == '}' || c == '[' || c == ']' || c == '(' || c == ')' {
				new += string(c)
			}
		}
		if new == "" {
			fmt.Println("OK")
			continue
		}
		count := 1
		for count > 0 {
			count = 0

			for i := 0; i < len(new)-1; i++ {
				if (new[i] == '{' && new[i+1] == '}') || (new[i] == '[' && new[i+1] == ']') || (new[i] == '(' && new[i+1] == ')') {
					new = new[:i] + new[i+2:]
					count++
					break
				}
			}
		}
		if new == "" {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
