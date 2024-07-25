package main

import "fmt"

func AddFront(s string, slice []string) []string {
	if s == "" {
		return slice
	}
	newSlice := make([]string, 0, len(slice)+len(s))
	// result := []string{}
	for _, str := range slice {
		newSlice = append(newSlice, s+str)
	}
	return newSlice
}

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}
