// given slice of ints and N write a pg that splits the slice if ints into N sizes.
// return slice of slice of ints where each slice of int is of size N
package main

import (
	"fmt"
)

func main() {
	// if len(os.Args) < 1 {
	// 	return
	// }
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := SplitSlice(a, 3)
	fmt.Println(r)
}

func SplitSlice(s []int, size int) [][]int {
	res := [][]int{}

	for i := 0; i < len(s); i += size {
		endSlice := i + size
		if endSlice >= len(s) {
			endSlice = len(s)
		}
		res = append(res, s[i:endSlice])
	}
	result := [][]int{}
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, res[i])
		// fmt.Println(result)
	}
	r := []int{}
	for _, vals := range result {
		for i := 0; i < len(vals); i++ {
			r = append(r, vals[i])
		}
	}
	fmt.Println(r)
	return result
}
