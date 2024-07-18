package main

import "fmt"

func AddIfPositive(a int, b int) int {
	// num := 0
	// if a > 0 && b > 0 {
	// 	num = a + b
	// } else {
	// 	return 0
	// }
	// return num
	val1 := a
	val2 := b

	num := 0

	if val1 > 0 && val2 > 0 {
		num = val1 + val2
	} else {
		return 0
	}
	return num
}

func main() {
	fmt.Println(AddIfPositive(1, 2))
	fmt.Println(AddIfPositive(1, -2))
	fmt.Println(AddIfPositive(-1, 2))
	fmt.Println(AddIfPositive(-1, -2))
	fmt.Println(AddIfPositive(10, 20))
	fmt.Println(AddIfPositive(0, 20))
}
