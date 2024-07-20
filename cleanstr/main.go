package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	var result string
	str := os.Args[1]
	if len(str) == 0 {
		fmt.Println()
		return
	}
	arr := []string{}
	for i, char := range str {
		if char != ' ' {
			result += string(char)
		}
		if (char == ' ' && result != "") || (i == len(str)-1 && result != "") {
			arr = append(arr, result)
			result = ""
		}

	}
	for i, v := range arr {
		if i != len(arr)-1 {
			fmt.Print(v + " ")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}
