package main

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

	for firstIn := 0; firstIn < len(slice); firstIn += size {
		lastInd := firstIn + size
		if lastInd > len(slice) {
			lastInd = len(slice)
		}
		result = append(result, slice[firstIn:lastInd])
	}
	fmt.Println(result)
}
func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 190, 2666666666666666666, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 51, 3255555, 3, 4, 588999, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3899, 4, 5555555555555555555, 6, 7}, 4)
}
