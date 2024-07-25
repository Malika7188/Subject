package main

import rvs "rsv"

func main() {
	rvs.Chunk([]int{}, 10)
	rvs.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	rvs.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	rvs.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	rvs.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
