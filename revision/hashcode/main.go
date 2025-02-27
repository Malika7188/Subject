package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func HashCode(dec string) string {
	size := len(dec)
	hash := []rune{}

	for _, c := range dec {
		hashed := (int(c) + size) % 127

		if hashed < 32 || hashed > 127 {
			hashed += 33
		}
		hash = append(hash, rune(hashed))
	}
	return string(hash)
}

func main() {
	table := []string{"Z", "Hi!", "BB198365", "sabito", "14 Avril 1912", "zyx987bca", "		pool-2020", "965truma747", " Mercedes-AMG GT"}
	for _, s := range table {
		challenge.Function("HashCode", HashCode, solutions.HashCode, s)
	}
}
