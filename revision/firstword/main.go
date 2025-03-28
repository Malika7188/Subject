package main

import (
	"fmt"
	"os"
)

var testCases = []struct {
	in   string
	want string
}{
	{"", "\n"},
	{"             a as", "a\n"},
	{"   f     d", "f\n"},
	{"   asd    ad", "asd\n"},
	{"   salut !!! ", "salut\n"},
	{" salut ! ! !", "salut\n"},
	{"salut ! !", "salut\n"},
}

func main() {
	for _, tc := range testCases {
		got := FirstWord(tc.in)
		if got != tc.want {
			fmt.Printf("FirstWord(%q) = %q instead of %q\n", tc.in, got, tc.want)
			os.Exit(1)
		}
	}
}

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}
	word := false
	res := ""

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			res += string(s[i])
			word = true
		} else if word {
			break
		}
	}
	return res + "\n"
}

// func main() {
// 	fmt.Print(FirstWord("hello there"))
// 	fmt.Print(FirstWord(""))
// 	fmt.Print(FirstWord("hello   .........  bye"))
// }
