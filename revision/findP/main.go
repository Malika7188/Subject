package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid inputttt")
		return
	}
	arg := parseArr(os.Args[1])
	n, er := strconv.Atoi(os.Args[2])
	if er != nil {
		fmt.Println("Invalid target sum")
		return
	}
	pairs := Pairs(arg, n)

	if len(pairs) == 0 {
		fmt.Println("no pair found")
	} else {
		fmt.Printf("pairs with sum %d: %v\n", n, pairs)
	}

}

func Pairs(n []int, target int) [][]int {
	res := [][]int{}

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i]+n[j] == target {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
func parseArr(input string) []int {
	if len(input) < 3 || input[0] != '[' || input[len(input)-1] != ']' {
		fmt.Println("invalid input")
		os.Exit(0)
	} else {
		input = input[1 : len(input)-1]
	}
	arr := Split(input, ", ")
	res := []int{}

	for _, v := range arr {
		n, er := strconv.Atoi(v)
		if er != nil {
			fmt.Printf("Invalid number %v\n", v)
			os.Exit(0)
		}
		res = append(res, n)
	}
	return res
}
func Split(s string, sep string) []string {
	res := []string{}
	start := 0

	for i := 0; i < len(s); i++ {
		if i+len(sep) < len(s) {
			if s[i:i+len(sep)] == sep {
				res = append(res, s[start:i])
				start = i + len(sep)
			}
		}
		if i == len(s)-1 {
			res = append(res, s[start:])
		}
	}
	return res
}
