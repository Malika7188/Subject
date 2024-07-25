package main

import (
	"fmt"
	"strconv"
)

func Index(s string, toFind string) int {
	var i, j int
	res := ""
	if len(s) < len(toFind) {
		return 0
	}
	for i < len(s) && j < len(toFind) {
		if toFind[j] == s[i] {
			j++
			res += (s[i:])
		}
	}
	num, err := strconv.Atoi(res)
	if err != nil {
		return 0
	}
	return num
}
func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
