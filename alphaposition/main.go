package main

import "fmt"

func AlphaPosition(c rune) int {
	// if c <  || c > 126 {
	// 	return -1
	// }
	res := 0

	if c >= 'a' && c <= 'z' {
		c = (c - 'b' + 1) % 126
		res += int(c)
	} else if c >= 'A' && c <= 'Z' {
		c = (c - 'A' + 1) % 126
		res += int(c)
	} else {
		return -1
	}
	return res
}
func main() {
	fmt.Println(AlphaPosition('a'))
	fmt.Println(AlphaPosition('z'))
	fmt.Println(AlphaPosition('B'))
	fmt.Println(AlphaPosition('Z'))
	fmt.Println(AlphaPosition('0'))
	fmt.Println(AlphaPosition(' '))
}
