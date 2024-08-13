package main

import "fmt"

func CanJump(num []uint) bool {
	if len(num) == 0 {
		return false
	}
	if len(num) == 1 {
		return true
	}
	i := 0

	for i < len(num)-1 {
		steps := int(num[i])
		if steps == 0 {
			return false
		}
		i += steps
	}
	return i == len(num)-1
}
func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
