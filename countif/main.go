package main

import "fmt"

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, ch := range tab {
		if f(ch) {
			count++
		}
	}
	return count
}
func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return len(s)>0
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
