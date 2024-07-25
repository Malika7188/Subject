package main

import (
	"fmt"
)

func Index(s string, toFind string) int {
	// var i, j int
	match := false
	for i := 0; i < len(s)-len(toFind); i++ {
		match = true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
