package main

import "fmt"

func AtoiBase(s string, base string) int {
	if len(base) < 1 {
		return 0
	}
	baseMap := make(map[rune]int)
	for i, c := range base {
		if c == '-' || c == '+' || baseMap[c] > 0 {
			return 0
		}
		baseMap[c] = i + 1

	}
	baselen := len(base)
	res := 0
	for _, c := range s {
		val, seen := baseMap[c]
		if !seen {
			return 0
		}
		res = res*baselen + (val - 1)
	}
	return res
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
