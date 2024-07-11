package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]

	num := Atoi(args)

	result := ""

	for i := 1; i <= 9; i++ {
		ans := i*num
		result += Itoa(i) + " * " + Itoa(num) + " = " + Itoa(ans) + "\n"
	}
	Printstr(result)
}
func Printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	//z01.PrintRune('\n')
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
	result := 0
	for _, val := range str {
		if val < '0' || val > '9' {
			os.Exit(0)
		}
		result = (result * 10) + int(val-'0')
	}
	if isneg == true {
		result *= -1
	}
	return result
}
func Itoa(number int) string {
	isneg := false
	if number < 0 {
		isneg = true
		number *= -1
	}

	result := ""
	if number == 0 {
		return "0"
	}
	for number > 0 {
		digit := number % 10
		result = string(digit+'0') + result
		number /= 10
	}
	if isneg == true {
		return "-" + result
	}
	return result
}