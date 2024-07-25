package main

import "fmt"

func HalfSlice(slice []int) []int {
	//result := []int{}

	LenSlice := len(slice)
	halslice := (LenSlice+1)/2

	return slice[:halslice]

}
func main() {
	fmt.Println(HalfSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(HalfSlice([]int{1, 2, 3}))
}
