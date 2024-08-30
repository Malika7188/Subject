package main

import "fmt"

func SumArray(numbers []int) int {
	res := 0
	if len(numbers) == 0 {
		return 0
	}

	for _, i := range numbers {
		res += i
	}
	return res
}
func main() {
	fmt.Println(SumArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(SumArray([]int{}))
	fmt.Println(SumArray([]int{-1, -2, -3, -4, -5}))
	fmt.Println(SumArray([]int{-1, 2, 3, 4, -5}))
}
