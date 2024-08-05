package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return -1
	}
	uniqueCh := make(map[rune]int)
	for _, ch := range str1 {
		uniqueCh[ch]++
	}
	for _, ch := range str2 {
		uniqueCh[ch]++
	}
	count := 0

	for _, ch := range uniqueCh {
		if ch == 1 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
