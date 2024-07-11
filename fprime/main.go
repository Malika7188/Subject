package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func IsPrime(num int) bool {
	for i := 2; i * i< num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Primefactors(number int) []int {
	result := []int{}
	for i := 2; i <= number; {
		if IsPrime(i) && number%i == 0 {
			result = append(result, i)
			number /= i

		} else if number%i != 0 {
			i++
		}
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	number, err := strconv.Atoi(args)
	if err != nil || number <= 0{
		return
	}

	factors := Primefactors(number)
	res := ""
	for _, val := range factors {
		res += strconv.Itoa(val) + "*"
	}
	res = res[:len(res)-1]

	for _, char := range res {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
