package rvs

import (
	"fmt"

	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	var result [][]int

	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	for firstI := 0; firstI < len(slice); firstI += size {
		lastI := firstI + size
		if lastI > len(slice) {
			lastI = len(slice)
		}
		result = append(result, slice[firstI:lastI])
	}
	fmt.Println(result)
}
