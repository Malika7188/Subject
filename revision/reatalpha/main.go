package main

import (
	"fmt"
	"os"
)

func RepeatAlpha(s string) string {
	res := ""
	repeat := 0
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			repeat = int(c - 'a' + 1)
		} else if c >= 'A' && c <= 'Z' {
			repeat = int(c - 'A' + 1)
		} else {
			repeat = 1
		}
		for i := 0; i < repeat; i++ {
			res += string(c)
			// z01.PrintRune(c)
		}
	}
	return res
}
func main() {
	testCases := []struct {
		in   string
		want string
	}{
		{"abc", "abbccc"},
		{"Choumi.", "CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii."},
		{"", ""},
		{"abacadaba 01!", "abbacccaddddabba 01!"},
	}
	for _, tc := range testCases {
		got := RepeatAlpha(tc.in)
		if got != tc.want {
			fmt.Printf("RepeatAlpha(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}
