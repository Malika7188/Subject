package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]
	var sum int
	//var newChar rune
	for _, char := range str {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			newChar := str[0]
			sum = int(rune(newChar) + char)
		}
	}
	fmt.Println(sum)
}