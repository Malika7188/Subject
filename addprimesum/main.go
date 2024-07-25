package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]

	numb, err := strconv.Atoi(args)
	if err != nil || numb < 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	res := 0
	for i := 2; i <= numb; i++ {
		if isprime(i) {
			res += i
		}
	}
	fmt.Println(res)
}

// func PrSum(num int) int {
// 	res := 0
// 	for i := 2; i <= num; i++ {
// 		if isprime(i) {
// 			res += i
// 		}
// 	}
// 	return res
//}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
