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
		if ch < '0' || ch > '9' {
			if ch != '-' && ch != '.' {
				return dec + "\n"
			}
		}
		if ch != '.' {
			res += string(ch)
		}
	}
	for res[0] == '0' && len(res) > 1 {
		res = res[1:]
	}
	if res == "0" {
		return "\n"
	}
	return res + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.001"))
	fmt.Print(NotDecimal("0.0000000000"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
