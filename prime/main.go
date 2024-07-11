package main

import (
// 	"fmt"
// 	"os"

// 	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	args := os.Args[1]

// 	num := Atoi(args)
// 	primefactors := isprime(num)

// 	if len(primefactors) == 0 {
// 		return
// 	}
// 	for i, factor := range primefactors {
// 		for i > 0 {
// 			z01.PrintRune('*')
// 		}
// 		numb := Itoa(factor)
// 		fmt.Println(numb)
// 	}
// }
func isprime(nb int) bool {
	for i := 2; i < nb; i++ {
		if i%2 == 0 {
			return false
		}
	}
	return true
}
func Atoi(str string) int {
	isneg := false
	if str[0] == '-' {
		isneg = true
		str = str[1:]
	}
	if str[0] == '+' {
		isneg = false
		str = str[1:]
	}
	var result int

	for _, char := range str {
		if char >= '0' && char <= '9' {
			result = (result * 10) + int(char-'0')
		} else {
			return 0
		}
	}
	if isneg {
		result *= -1
	}
	return result
}
func Itoa(nb int) string {
	isneg := false
	if nb < 0 {
		isneg = true
		nb *= -1
	}
	result := "" 

	for nb > 0 {
		digit := nb%10
		result = string(digit + '0') + result
		nb /= 10
	} 
	if isneg == true {
		return result + "-"
	}
	return result
}