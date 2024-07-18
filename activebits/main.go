package main

import "fmt"

func ActiveBits(n int) int {
	var bit int

	for n > 0 {
		digit := n%2
		bit = digit + bit
		n /= 2
	}
	return bit


}

func main() {
	fmt.Println(ActiveBits(8))
}
