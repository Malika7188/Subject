package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return-1
	}
	seen1 := make(map[rune]bool)
	seen2 := make(map[rune]bool)

	for _, ch := range str1 {
		seen1[ch] = true
	}
	for _, ch := range str2 {
		seen2[ch] = true
	}
	count := 0
	for ch := range seen1 {
		if !seen2[ch] {
			count++
		}
	}
	for ch := range seen2 {
		if !seen1[ch] {
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
