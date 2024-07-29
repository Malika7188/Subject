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
	fmt.Println(SplitSlice([]int{0, 1, 2, 3, 4, 5, 6, 7,8,9}, 3))
}

func SplitSlice(s []int, size int) [][]int {
	res := [][]int{}
	N := size
	for i := 0; i < len(s); i += N {
		endSlice := i + N
		if endSlice > len(s) {
			endSlice = len(s)
		}
		res = append(res, s[i:endSlice])
	}
	return res
}
