package revision

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

	for firstIn := 0; firstIn < len(slice); firstIn+= size {
		lastInd := firstIn + size 
		if lastInd > len(slice) {
			lastInd = len(slice)
		}
		result = append(result, slice[firstIn:lastInd])
	}
	fmt.Println(result)
}
 