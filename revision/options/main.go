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
	arg := os.Args[1:]

	var option int32

	for _, str := range arg {
		if str == "" || str == "-" {
			fmt.Println("Invalid Option")
			return
		}

		for i, c := range str[1:] {
			if !Valid(c) {
				fmt.Println("Invalid Option")
				return
			}
			if i == 0 && c == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			option |= 1 << (c - 'a')
		}

	}
	fmt.Printf("%08b, %08b, %08b,%08b", (option>>24)&0xff, (option>>16)&0xff, (option>>8)&0xff, (option)&0xff)
	fmt.Println()
}
func Valid(s rune) bool {
	return s >= 'a' && s <= 'z'
}
