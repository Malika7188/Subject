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
	new := ""
	for _, str := range args {
		for _, ch := range str {
			if ch == '{' || ch == '}' || ch == '[' || ch == ']' || ch == '(' || ch == ')' {
				new += string(ch)
			}
		}
		// fmt.Println(new)
		count := 1
		for count > 0 && len(new) > 2 {
			count = 0
			for i := 0; i < len(new); i++ {
				if new[i] == '{' && new[i+1] == '}' {
					new = new[:i] + new[i+2:]
					count++
				}
				if new[i] == '[' && new[i+1] == ']' {
					new = new[:i] + new[i+2:]
					count++
				}
				if new[i] == '(' && new[i+1] == ')' {
					new = new[:i] + new[i+2:]
					count++
				}

			}

		}
		if new == "{}" || new == "[]" || new == "()" || str == ""{
			fmt.Println("OK")
		} else {
			fmt.Println("ERROR")
		}
	}
}
