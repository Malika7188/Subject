package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	arg := strings.Fields(os.Args[1])

	num := []int{}

	for _, c := range arg {
		if c == "+" || c == "-" || c == "*" || c == "/" || c == "%" {
			if len(num) < 2 {
				fmt.Println("Error")
			}
			a := num[len(num)-2]
			b := num[len(num)-1]
			num = num[:len(num)-2]
			cal := 0

			if c == "+" {
				cal = a + b
			} else if c == "-" {
				cal = a - b
			} else if c == "/" {
				cal = a / b
			} else if c == "*" {
				cal = a * b
			} else if c == "%" {
				cal = a % b
			} else {
				fmt.Println("invalid input")
				return
			}
			num = append(num, cal)
		} else {
			n, er := strconv.Atoi(c)
			if er != nil {
				fmt.Println("Error")
				return
			}
			num = append(num, n)
		}
	}
	if len(num) == 1 {
		fmt.Println(num[0])
	} else {
		fmt.Println("Error")
	}

}

// func Split(s, sep string) []string {
// 	res := []string{}
// 	start := 0

// 	for i := 0; i < len(s); i++ {
// 		if i+len(sep) < len(s) {
// 			if s[i:i+len(sep)] == sep {
// 				res = append(res, s[start:i])
// 				start = i + len(sep)
// 			}
// 		}
// 		if i == len(s)-1 {
// 			res = append(res, s[start:])
// 		}
// 	}
// 	return res
// }
