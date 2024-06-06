package main

import (
	"fmt"
	"os"
	
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	
	str1 := args[0]
	str2 := args[1]
	result := ""

for _, char := range str1 {
		output := false 
		for i, char1 := range str2 {
			if char == char1 {
				result += string(char1)
				str2 = str2[i:]

				output = true
				break
			}
		}
		if !output {
			return
		}
		// for _, char := range result {
		// 	z01.PrintRune(char)
		// }
	}
	fmt.Println(result)


}