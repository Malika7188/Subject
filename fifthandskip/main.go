package main

import (
	"fmt"
	"os"
)

func FifthAndSkip(str string) string {
	if len(str) == 0 {
		return ""
	}
	word := ""
	for _, char := range str {
		if char == ' ' {
			continue
		}
		word += string(char)
	}
	if len(word) < 5 {
		os.Stdout.WriteString("invalid input\n")
		return ""
	}
	result := ""
	for i := 0; i < len(word); i += 5 {
		if i+5 > len(word) {
			result += word[i:]
		} else {
			result += word[i:i+5] + " "
		}
		i++
	}
	return result + "\n"
}
func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
	fmt.Print(FifthAndSkip(""))
}
