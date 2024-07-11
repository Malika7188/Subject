package main

import "revision"

func main() {
	revision.Chunk([]int{}, 10)
	revision.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	revision.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	revision.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	revision.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
