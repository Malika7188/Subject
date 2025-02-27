package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	// if str1 
	seen1 := make(map[rune]bool)
	seen2 := make(map[rune]bool)

	for _, c := range str1 {
		seen1[c] = true
	}
	for _, ch := range str2 {
		seen2[ch] = true
	}
	count := 0
	for cs := range seen1 {
		if !seen2[cs] {
			count++
		}
	}
	for char := range seen2 {
		if !seen1[char] {
			count++
		}
	}
	return count
}
func main() {
	table := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
	}
	for _, arg := range table {
		challenge.Function("WeAreUnique", WeAreUnique, solutions.WeAreUnique, arg[0], arg[1])
	}
}
