package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	if s == "" {
		return s
	}
	count := 1
	res := ""
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			res += strconv.Itoa(count) + string(s[i-1])
			count = 1
		}
	}
	res += strconv.Itoa(count) + string(s[len(s)-1])
	return res
}
func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
