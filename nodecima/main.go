package main

import (
	"fmt"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	if len(dec) == 1 {
		return dec
	}
	res := ""
	for _, ch := range dec {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			return dec
		}
		if ch != '.' {
			res += string(ch)
		}
	}
	for res[0] == '0' {
		res = res[1:]
	}
	return res + "\n"
}

// func treamS(str string) string {
// 	res := ""
// 	for i, ch := range str {
// 		if ch == '0' && i == 0 {
// 			res += str[1:]
// 		}
// 	}
// 	return res
// }

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
