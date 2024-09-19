package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	args := os.Args[1:]

	slice := args[0]
	slice2 := Split(args[1], " ")

	if len(slice) == 0 {
		return
	}
	if len(slice) < 2 || slice[0] != '(' || slice[len(slice)-1] != ')' {
		return
	} else {
		slice = slice[1 : len(slice)-1]
	}

	slice1 := Split(slice, "|")
	count := 0
	for _, str := range slice2 {
		for _, s := range slice1 {
			if StringContains(str, s) {
				count++
				if str[len(str)-1] == ',' {
					str = str[:len(str)-1]
				}
				fmt.Printf("%v: %v\n", count, str)
			}
		}
	}
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

func StringContains(s string, r string) bool {
	if len(r) == 0 || len(s) == 0 {
		return false
	}
	for i := 0; i <= len(s)-len(r); i++ {
		if s[i:i+len(r)] == r {
			return true
		}
	}

	return false
}
