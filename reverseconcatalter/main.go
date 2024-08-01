package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	newSlice := []int{}
	len1, len2 := len(slice1), len(slice2)

	i, j := len1-1, len2-1

	for i >= 0 && j >= 0 {
		if len1 > len2 {
			newSlice = append(newSlice, slice1[i])
			len1--
			i--
		} else if len1 < len2 {
			newSlice = append(newSlice, slice2[j])
			len2--
			j--
		} else {
			newSlice = append(newSlice, slice1[i])
			len1--
			i--
		}
	}
	for i >= 0 {
		newSlice = append(newSlice, slice1[i])
		i--
	}
	for j >= 0 {
		newSlice = append(newSlice, slice2[j])
		j--
	}
	return newSlice
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
	fmt.Println(RevConcatAlternate([]int{}, []int{}))

}
