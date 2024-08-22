package main

import (
	"fmt"
)

func MergeAndSort(a, b []int) []int {

	for i := 0; i < len(b); i++ {
		a = append(a, b[i])
	}
	// a = a[len(a)-1]
	return sort(a)
}
func main() {
	fmt.Println(MergeAndSort([]int{1, 4, 2, 3}, []int{9, 8, 5, 6, 10, 7}))
}
func sort(a []int) []int{
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a)-1; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
