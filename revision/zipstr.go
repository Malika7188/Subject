package main

import (
	"strconv"
)

func ZipString(s string) string {
	if s == "" {
		return s
	}
	res := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			res += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
		// count = 0
	}
	res += strconv.Itoa(count) + string(s[len(s)-1])
	return res
}

// func main() {
// 	args := []string{
// 		"aaaa",
// 		"zzzzzZZZZZZ",
// 		"",
// 		"ziiiiipMeee",
// 		"hhellloTthereYouuunggFelllas",
// 	}

// 	for _, arg := range args {
// 		fmt.Printf("ZipString(%q) = %q\n", arg, ZipString(arg))
// 	}
// }
