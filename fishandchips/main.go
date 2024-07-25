package main

import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}
	for i := 0; i <= n; i++ {
		if n%6 == 0 {
			return "fish and chips"
		}
		if n%2 == 0 {
			return "fish"
		}
		if n%3 == 0 {
			return "chips"
		}
	}
	return "fishandchips"
}
func main() {
	fmt.Println(FishAndChips(4))
	fmt.Println(FishAndChips(9))
	fmt.Println(FishAndChips(6))
}
