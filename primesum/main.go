package main

import (
	"fmt"
	"os"
	"strconv"

	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	args := os.Args[1]

	number, err := strconv.Atoi(args)
	if err != nil || number < 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(SumOfPrime(number))
}

func IsPrime(num int) bool {
	if num < 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func SumOfPrime(number int) int {
	total := 0

	for i := 2; i <= number; i++ {
		if IsPrime(i) {
			total += i
		}
	}
	return total
}
