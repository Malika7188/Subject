package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	s3 := os.Args[3]
	res := ""
	// seen := false
	for _, c := range s1 {
		for _, ch := range s2 {
			if c == ch {
				res += s3
			} else {
				res += string(c)
			}
		}
	}
	fmt.Println(res)
}