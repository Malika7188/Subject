package main

import (
	"fmt"
	// "sort"
)

func Abort(a, b, c, d, e, f,g int) int {
	num := []int{a, b, c, d, e, f,g}
	mid := len(num)
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	//sort.Ints(num)
	if mid%2 != 0 {
		return num[mid/2]
	} else {
		return (num[mid/2-1] + num[mid/2]) / 2
	}
}

func main() {
	middle := Abort(2, 3, 8, 5, 7, 9,10)
	fmt.Println(middle)
}
