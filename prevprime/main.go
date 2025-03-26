package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

// import "fmt"

func FindPrevPrime(nb int) int {
	if Prime(nb) {
		return nb
	} else {
		for i := nb; i >= 2; i-- {
			if Prime(i) {
				return i
			}
		}
	}
	return 0
}

func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	args := append(random.IntSliceBetween(0, 99999), 5, 4, 1, 0)
	for _, arg := range args {
		challenge.Function("FindPrevPrime", FindPrevPrime, solutions.FindPrevPrime, arg)
	}
}
