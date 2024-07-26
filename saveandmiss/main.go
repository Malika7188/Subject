package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	res := ""
	if num <= 0 {
		return arg
	}

	for i := 0; i < len(arg); i+=num*2 {
		end := i + num
		if end > len(arg) {
			end = len(arg)
		}
		res += string(arg[i:end])
	}
	return res
}
func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
