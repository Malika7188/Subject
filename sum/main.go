package main

import (
	"fmt"
	"strconv"
)

func Sum(a, b string) int {
	num1, err := strconv.Atoi(a)
	CheckErr(err)
	num2, err := strconv.Atoi(b)
	CheckErr(err)

	sum := num1 + num2
	//strconv.Itoa(sum)
	return sum
}

func CheckErr(err error) {
	if err != nil {
		return
	}
}

func main() {
	fmt.Println(Sum("1", "3"))
	fmt.Println(Sum("-4", "1"))
	fmt.Println(Sum("7", "-5"))
}
