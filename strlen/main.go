package main

import (
	"fmt"
	//"github.com/01-edu/z01"
	//"piscine"
)

func StrLen(s string) int {
	// return len([]rune(s))
	m := []rune(s)
	return len(m)
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
