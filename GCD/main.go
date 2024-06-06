package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	a := Atoi(os.Args[1])
	b := Atoi(os.Args[2])

	gcd := GCD(a, b)

	fmt.Println(gcd)
}

func Atoi(str string) int {
	var isneg bool
	if str[0] == '-' {
		str = str[1:]
		isneg = true
	}
	if str[0] == '+' {
		str = str[1:]
	}
	var num int
	for _, val := range str {
		if val >= '0' && val <= '9' {
			num = (num * 10) + int(val-'0')
		} else {
			return 0
		}
	}
	if isneg {
		num *= -1
	}
	return num
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
