package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	arg := os.Args[1]

	n, err := strconv.Atoi(arg)
	if err != nil {
		return
	}
	// if n <= 1 {
	// 	z01.PrintRune('0')
	// }
	res := Psum(n)

	fmt.Println(res)

}
func Psum(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			res += i
		}
	}
	return res
}
func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
