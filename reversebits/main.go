package main

import (
	"fmt"
)

func ReverseBits(oct byte) byte {
	var reversed byte = 0
	for i := 0; i < 8; i++ {
		reversed = (reversed << 1) | (oct & 1)
		oct >>= 1
	}
	return reversed
}

func main() {
	oct := byte(0b00100110) // Example byte
	reversed := ReverseBits(oct)
	fmt.Printf("%08b || / %08b\n", oct, reversed)
}
