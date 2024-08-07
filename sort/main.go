package main

import "fmt"

func BubSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
func main() {
	num := []int{27,76,4,33,65}
	fmt.Println(BubSort(num))
}
