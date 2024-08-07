package main

import "fmt"

func AppendRange(min, max int) []int {
	if min >= max {
		return []int{}
	}
	result := []int{}
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
