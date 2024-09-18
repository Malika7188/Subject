package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func ItoaBase(num, b int) string {
	if num == 0 {
		return "0"
	}
	hexv := "0123456789ABCDEF"
	res := ""
	sign := ""
	if num < 0 {
		num *= -1
		sign = "-"
	}
	value := uint(num)
	base := uint(b)
	for value > 0 {
		dig := value % base
		res = string(hexv[dig]) + res
		value /= base
	}
	return sign + res
}

// func main() {
// 	fmt.Println(ItoaBase(10, 2))
// 	fmt.Println(ItoaBase(255, 16))
// 	fmt.Println(ItoaBase(-42, 4))
// 	fmt.Println(ItoaBase(123, 10))
// 	fmt.Println(ItoaBase(0, 8))
// 	fmt.Println(ItoaBase(255, 2))
// 	fmt.Println(ItoaBase(-255, 16))
// 	fmt.Println(ItoaBase(15, 16))
// 	fmt.Println(ItoaBase(10, 4))
// 	fmt.Println(ItoaBase(255, 10))
// }

func main() {
	for i := 0; i < 30; i++ {
		value := random.IntBetween(-1000000, 1000000)
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, value, base)
	}
	for i := 0; i < 5; i++ {
		base := random.IntBetween(2, 16)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MaxInt, base)
		challenge.Function("ItoaBase", ItoaBase, solutions.ItoaBase, random.MinInt, base)
	}
}
