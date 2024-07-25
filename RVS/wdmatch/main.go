 package main

import (
	"github.com/01-edu/z01"
  "os"
)
func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	if wdmatch(str1, str2) {
		for _, char := range str1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}
func wdmatch(str1, str2 string) bool {
	var i, j int
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++
	}
	if i == len(str1) {
		return true
	} else {
		return false
	}
}
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 3 {
// 		return
// 	}

// 	str1 := args[0]
// 	str2 := args[1]

// 	if wdmatch(str1, str2) {
// 		for _, char := range str1 {
// 			z01.PrintRune(char)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		return
// 	}
	

// }
// func wdmatch(str1, str2 string) bool{
// 	var i,j int

// 	for i < len(str1) && j < len(str2) {
// 		if str1[i] == str2[j] {
// 			i++
// 		}
// 		j++
// 	}
// 	if i == len(str1) {
// 		return true
// 	} else {
// 		return false
// 	}
// }
