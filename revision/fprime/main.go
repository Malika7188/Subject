package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	arg := os.Args[1]

	n, err := strconv.Atoi(arg)
	if err != nil || n <= 1 {
		return
	}
	res := Fact(n)
	ans := ""
	for _, c := range res {
		ans += strconv.Itoa(c) + "*"
	}
	ans = ans[:len(ans)-1]
	fmt.Println(ans)
}
func Fact(n int) []int {
	res := []int{}

	for i := 2; i <= n; {
		for IsPrime(i) && n%i == 0 {
			res = append(res, i)
			n /= i
		}
		i++

	}
	return res
}
func IsPrime(n int) bool {
	for i := 2; i *i< n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
