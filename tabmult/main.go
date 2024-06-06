package main

import (
	"fmt"
	"os"
	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	num := Atoi(os.Args[1])

	for i := 1; i <= 9; i++ {
		fmt.Printf("%v * %v = %v\n", i, num, i*num)
	}
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
