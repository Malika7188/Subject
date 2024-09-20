package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func FifthAndSkip(str string) string {

	if str == "" {
		return "\n"
	}
	// nonspaces := 0
	// for _, c := range str {
	// 	if c != ' ' {
	// 		nonspaces++
	// 	}
	// }
	if len(str) < 5 {
		return ("Invalid Input\n")

	}
	keep := []rune{}
	count := 0
	for _, c := range str {
		if c == ' ' {
			continue
		}
		if count == 5 {
			keep = append(keep, ' ')
			count = 0
			continue
		}
		count++
		keep = append(keep, c)
	}
	return string(keep) + "\n"

}

// func main() {
// 	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
// 	fmt.Print(FifthAndSkip("This is a short sentence"))
// 	fmt.Print(FifthAndSkip("1234"))
// 	fmt.Print(FifthAndSkip(""))
// 	fmt.Print(FifthAndSkip("e 5£jhy @ a* 7 =56 ;"))
// }

func main() {
	table := []string{"1234556789", "e 5£ jhy@ a* 7 =56 ;", "QKplq%QSw", "", "hello \\! n4ght cr3a8ure7 ", "Kimetsu no Yaiba", "8595485-52", "-552", "w58tw7474abc", "Po65 4o"}
	for _, s := range table {
		challenge.Function("FifthAndSkip", FifthAndSkip, solutions.FifthAndSkip, s)
	}
}
