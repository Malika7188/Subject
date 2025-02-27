package main

import "fmt"

func SaveAndMiss(arg string, num int) string {
	if num < 1 {
		return arg
	}
	res := ""
	for i := 0; i < len(arg); i++ {
		if (i/num)%2 == 0 {
			res += string(arg[i])
		}
	}
	return res
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello \\! n4ght cr3a8ure7 ", 5))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}

// func main() {
// 	fmt.Println(SaveAndMiss("e 5£ @ 8* 7 =56 ;", 2))
// 	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
// 	fmt.Println(SaveAndMiss("", 3))
// 	fmt.Println(SaveAndMiss("hello you all ! ", 0))
// 	fmt.Println(SaveAndMiss("what is your name?", 0))
// 	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))

// }

// func main() {
// 	structs := []struct {
// 		str string
// 		itt int
// 	}{
// 		{"123456789", 1},
// 		{"e 5S @ 8* 7 =56 ;", 2},
// 		{"QKplq%QSw", 3},
// 		{"", 4},
// 		{"hello \\! n4ght cr3a8ure7 ", 5},
// 		{"Kimetsu no Yaiba", 6},
// 		{"8595485-52", 7},
// 		{"abcdefghijklmnopqrstuvwyz", 8},
// 		{"w58tw7474abc", 9},
// 		{"Po65 4o", 10},
// 	}

// 	// for _, v := range structs {
// 	// 	Function(SaveAndMiss(), SaveAndMiss, SaveAndMiss, v.str, v.itt)
// 	// }
// }
