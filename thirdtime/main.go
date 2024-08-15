package main

import "fmt"

func ThirdTimeIsACharm(str string) string {
	res := ""
	for i := 2; i < len(str); i += 3 {
		res += string(str[i])
	}
	return res + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
