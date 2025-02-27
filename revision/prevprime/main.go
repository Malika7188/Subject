package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

// import "fmt"

func FindPrevPrime(nb int) int {
	for nb > 1 && Prime(nb) {
		return nb
	}
	nb--
	if nb < 1 {
		return nb
	}

	return nb
}

func Prime(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}
func main() {
	args := append(random.IntSliceBetween(0, 99999), 5, 4, 1, 0)
	for _, arg := range args {
		challenge.Function("FindPrevPrime", FindPrevPrime, solutions.FindPrevPrime, arg)
	}
}
