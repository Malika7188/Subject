package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	args := os.Args[1:]

	var options int32
	for _, str := range args {
		if str == "" {
			fmt.Println("inavaalid options")
			return
		}
		for i, c := range str[1:] {
			if !valid(c) {
				fmt.Println("invalid options")
				return
			}
			if i == 0 && c == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			options |= 1 << (c - 'a')
		}
	}
	fmt.Printf("%08b %08b %08b %08b", (options>>24)&0xff, (options>>16)&0xff, (options>>8)&0xff, (options)&0xff)
	fmt.Println()
}
func valid(c rune) bool {
	return c >= 'a' && c <= 'z'
}
